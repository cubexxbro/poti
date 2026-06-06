package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

const asciiArt = `
          _   _ 
 ___  ___| |_(_)
| '_ \/ _ \ __| |
| |_) | (_) | |_| |
| .__/ \___/\__|_|
|_|              
`

const cacheDirName = ".poti_cache"

type ScanResult struct {
	IP     string
	Port   int
	Banner string
}

type ProgressProxy struct {
	Total   int64
	Current int64
}

func (pp *ProgressProxy) Write(p []byte) (int, error) {
	n := len(p)
	pp.Current += int64(n)
	if pp.Total > 0 {
		pct := (pp.Current * 100) / pp.Total
		fmt.Printf("\r\033[1;34m[*] Synchronizing global matrix tables... [%d%%]\033[0m", pct)
	} else {
		fmt.Printf("\r\033[1;34m[*] Synchronizing global matrix tables... [%d KB]\033[0m", pp.Current/1024)
	}
	return n, nil
}

func getCacheFilePath(cc string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		return filepath.Join(".", cacheDirName, cc+".txt")
	}
	return filepath.Join(home, cacheDirName, cc+".txt")
}

func ensureCacheDir() {
	home, err := os.UserHomeDir()
	var path string
	if err != nil {
		path = filepath.Join(".", cacheDirName)
	} else {
		path = filepath.Join(home, cacheDirName)
	}
	_ = os.MkdirAll(path, 0755)
}

func loadIPRanges(cc string) ([]string, error) {
	ensureCacheDir()
	cachePath := getCacheFilePath(cc)

	if _, err := os.Stat(cachePath); err == nil {
		return readLines(cachePath)
	}

	fmt.Printf("\033[1;34m[*] Resolving remote nodes for '%s'...\033[0m\n", cc)
	
	url := fmt.Sprintf("https://raw.githubusercontent.com/herrbischoff/country-ip-blocks/master/ipv4/%s.txt", strings.ToLower(cc))
	if cc == "CN" {
		url = "https://raw.githubusercontent.com/gaoyifan/china-operator-ip/ip-lists/china.txt"
	}

	client := http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("remote registry returned status: %d (Invalid code?)", resp.StatusCode)
	}

	totalBytes, _ := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)

	out, err := os.Create(cachePath)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	proxy := &ProgressProxy{Total: totalBytes}
	teeReader := io.TeeReader(resp.Body, proxy)

	_, err = io.Copy(out, teeReader)
	fmt.Println()
	if err != nil {
		return nil, err
	}

	return readLines(cachePath)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") && strings.Contains(line, "/") {
			lines = append(lines, line)
		}
	}
	if len(lines) == 0 {
		return nil, fmt.Errorf("empty asset pool: no valid CIDR segments detected")
	}
	return lines, scanner.Err()
}

func clearAllCache() {
	home, err := os.UserHomeDir()
	var path string
	if err != nil {
		path = filepath.Join(".", cacheDirName)
	} else {
		path = filepath.Join(home, cacheDirName)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("\033[1;32m[+] No downloaded data found. System is already clean.\033[0m")
		return
	}

	err = os.RemoveAll(path)
	if err != nil {
		fmt.Printf("\033[1;31m[-] Error cleaning data: %v\033[0m\n", err)
	} else {
		fmt.Println("\033[1;32m[+] [Success] Cache directory purged.\033[0m")
	}
}

func grabBanner(ip string, port int, timeout time.Duration) string {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return ""
	}
	defer conn.Close()

	_ = conn.SetDeadline(time.Now().Add(timeout))

	if port == 80 || port == 8080 || port == 443 {
		fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: %s\r\nUser-Agent: poti/0.7.2\r\nConnection: close\r\n\r\n", ip)
	}

	scanner := bufio.NewScanner(conn)
	var bannerLines []string
	
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" && (port == 80 || port == 8080 || port == 443) {
			break
		}
		if len(bannerLines) < 5 {
			bannerLines = append(bannerLines, "    │ "+strings.TrimSpace(line))
		} else {
			break
		}
	}

	if len(bannerLines) > 0 {
		return strings.Join(bannerLines, "\n")
	}
	return "    │ [No immediate response / Hidden Active Service]"
}

func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func getIPsFromCIDR(cidr string) []string {
	var ips []string
	ip, ipnet, err := net.ParseCIDR(strings.TrimSpace(cidr))
	if err != nil {
		return ips
	}

	for el := ip.Mask(ipnet.Mask); ipnet.Contains(el); incIP(el) {
		ips = append(ips, el.String())
	}
	
	if len(ips) > 2 {
		return ips[1 : len(ips)-1]
	}
	return ips
}

func pickRandomIPs(ipPool []string, count int) []string {
	if len(ipPool) == 0 {
		return nil
	}
	if count > len(ipPool) {
		count = len(ipPool)
	}

	result := make([]string, count)
	chosen := make(map[int]bool)

	for i := 0; i < count; i++ {
		for {
			nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(len(ipPool))))
			idx := int(nBig.Int64())
			if !chosen[idx] {
				chosen[idx] = true
				result[i] = ipPool[idx]
				break
			}
		}
	}
	return result
}

func worker(tasks <-chan string, ports []int, results chan<- ScanResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for ip := range tasks {
		for _, port := range ports {
			banner := grabBanner(ip, port, 2*time.Second)
			if banner != "" {
				results <- ScanResult{IP: ip, Port: port, Banner: banner}
			}
		}
	}
}

func main() {
	fmt.Printf("\033[1;36m%s\033[0m", asciiArt)
	fmt.Println("\033[1;32m[+] poti Engine - Global Intelligence Radar v0.7.2\033[0m")
	fmt.Println("--------------------------------------------------")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\033[1;33m[?] Enter Target Country Code (e.g. US, CN, JP, KR, DE) or 'clear':\033[0m")
	fmt.Print("Input (default CN): ")
	ccInput, _ := reader.ReadString('\n')
	ccInput = strings.TrimSpace(strings.ToUpper(ccInput))

	if ccInput == "CLEAR" {
		fmt.Println("\n[*] Running secure system cleaner...")
		clearAllCache()
		return
	}

	cc := "CN"
	if ccInput != "" {
		cc = ccInput
	}

	ranges, err := loadIPRanges(cc)
	if err != nil {
		fmt.Printf("\033[1;31m[-] Synchronization failure: %v\033[0m\n", err)
		return
	}

	fmt.Print("\n\033[1;33m[?] Enter extraction quota (How many random targets? default 10):\033[0m ")
	limitInput, _ := reader.ReadString('\n')
	limitInput = strings.TrimSpace(limitInput)
	limit := 10
	if limitInput != "" {
		fmt.Sscanf(limitInput, "%d", &limit)
	}
	if limit <= 0 {
		limit = 10
	}

	fmt.Print("\n\033[1;33m[?] Enter ports to scan (comma-separated, default 80,443,8080):\033[0m ")
	portsInput, _ := reader.ReadString('\n')
	portsInput = strings.TrimSpace(portsInput)
	if portsInput == "" {
		portsInput = "80,443,8080"
	}

	fmt.Println("\n[*] Formulating targeting lattice matrix from large scale assets...")
	var allIPs []string
	for _, cidr := range ranges {
		cidr = strings.TrimSpace(cidr)
		if cidr != "" && !strings.Contains(cidr, ":") {
			allIPs = append(allIPs, getIPsFromCIDR(cidr)...)
		}
	}

	if len(allIPs) == 0 {
		fmt.Println("\033[1;31m[-] Target pool generation yielded zero viable endpoints.\033[0m")
		return
	}

	targets := pickRandomIPs(allIPs, limit)
	
	var targetPorts []int
	portSpecs := strings.Split(portsInput, ",")
	for _, spec := range portSpecs {
		var p int
		fmt.Sscanf(strings.TrimSpace(spec), "%d", &p)
		if p > 0 && p <= 65535 {
			targetPorts = append(targetPorts, p)
		}
	}

	fmt.Println("--------------------------------------------------")
	fmt.Printf("[*] Launching Radar Matrix Mode...\n")
	fmt.Printf("[*] Target Region    : %s\n", cc)
	fmt.Printf("[*] Total Available Pool Size: %d live public subnets\n", len(allIPs))
	fmt.Printf("[*] Extracted Audit Targets  : %d random nodes\n", len(targets))
	fmt.Printf("[*] Target Verification Ports: %v\n", targetPorts)
	fmt.Println("--------------------------------------------------\n[*] Mapping live space assets across the world...")

	tasksChan := make(chan string, len(targets))
	resultsChan := make(chan ScanResult, 100)
	var wg sync.WaitGroup

	numWorkers := 30
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(tasksChan, targetPorts, resultsChan, &wg)
	}

	for _, ip := range targets {
		tasksChan <- ip
	}
	close(tasksChan)

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	found := 0
	for res := range resultsChan {
		found++
		fmt.Printf("\n [🔥] Found Exposed Asset: \033[1;33m%s\033[0m:\033[1;32m%d\033[0m\n", res.IP, res.Port)
		fmt.Println(res.Banner)
	}

	fmt.Println("\n--------------------------------------------------")
	fmt.Printf("[*] Real-world intelligence cycle complete. Total discoveries: %d\n", found)
}
