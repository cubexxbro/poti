package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"net"
	"os"
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

var countryIPRanges = map[string][]string{
	"CN": {"1.0.1.0/24", "1.0.2.0/23", "1.1.1.0/24", "14.116.0.0/16", "116.62.0.0/16"},
	"US": {"8.8.8.0/24", "13.107.21.0/24", "34.192.0.0/12", "104.16.0.0/12"},
	"JP": {"1.0.64.0/18", "1.1.64.0/24", "118.238.0.0/16"},
}

type ScanResult struct {
	IP     string
	Port   int
	Banner string
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
		fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: %s\r\nUser-Agent: poti/0.5.0\r\nConnection: close\r\n\r\n", ip)
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
	ip, ipnet, err := net.ParseCIDR(cidr)
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
	fmt.Println("\033[1;32m[+] poti Engine - Smart Interactive Radar v0.5.0\033[0m")
	fmt.Println("--------------------------------------------------")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\033[1;33m[?] Select Target Geographical Region:\033[0m")
	fmt.Println("  1. China (CN)")
	fmt.Println("  2. United States (US)")
	fmt.Println("  3. Japan (JP)")
	fmt.Print("Choose option (1-3, default 1): ")
	ccInput, _ := reader.ReadString('\n')
	ccInput = strings.TrimSpace(ccInput)

	cc := "CN"
	if ccInput == "2" {
		cc = "US"
	} else if ccInput == "3" {
		cc = "JP"
	}

	fmt.Print("\n\033[1;33m[?] How many random IPs to extract? (default 5):\033[0m ")
	limitInput, _ := reader.ReadString('\n')
	limitInput = strings.TrimSpace(limitInput)
	limit := 5
	if limitInput != "" {
		fmt.Sscanf(limitInput, "%d", &limit)
	}
	if limit <= 0 {
		limit = 5
	}

	fmt.Print("\n\033[1;33m[?] Enter ports to scan (comma-separated, default 80,443,8080):\033[0m ")
	portsInput, _ := reader.ReadString('\n')
	portsInput = strings.TrimSpace(portsInput)
	if portsInput == "" {
		portsInput = "80,443,8080"
	}

	ranges := countryIPRanges[cc]
	var allIPs []string
	for _, cidr := range ranges {
		allIPs = append(allIPs, getIPsFromCIDR(cidr)...)
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

	fmt.Println("\n--------------------------------------------------")
	fmt.Printf("[*] Launching Radar Mode...\n")
	fmt.Printf("[*] Target Region  : %s\n", cc)
	fmt.Printf("[*] Extracted Nodes: %d targets\n", len(targets))
	fmt.Printf("[*] Target Ports   : %v\n", targetPorts)
	fmt.Println("--------------------------------------------------\n[*] Mapping live space assets...")

	tasksChan := make(chan string, len(targets))
	resultsChan := make(chan ScanResult, 100)
	var wg sync.WaitGroup

	numWorkers := 10
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
	fmt.Printf("[*] Smart check complete. Total discoveries: %d\n", found)
}
