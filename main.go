package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println("[+] poti v0.7.5 - Global Turbo Engine")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Target CC: ")
	cc, _ := reader.ReadString('\n')
	cc = strings.TrimSpace(strings.ToUpper(cc))
	if cc == "" { cc = "CN" }

	fmt.Println("[*] Fetching global assets...")
	ips := loadGlobalAssets(cc)
	if len(ips) == 0 { fmt.Println("[-] No assets found."); return }

	fmt.Print("Quota: ")
	q, _ := reader.ReadString('\n')
	limit := 100
	fmt.Sscanf(strings.TrimSpace(q), "%d", &limit)

	fmt.Println("[*] Radar active...")
	results := make(chan string, limit)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, ip := range pickRandomIPs(ips, limit/100 + 1) {
				if check(ip, 80) { results <- ip }
			}
		}()
	}
	go func() { wg.Wait(); close(results) }()
	for res := range results { fmt.Printf("[!] Found: %s\n", res) }
}

func loadGlobalAssets(cc string) []string {
	cachePath := filepath.Join(os.TempDir(), "poti_"+cc+".cache")
	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
		url := fmt.Sprintf("https://raw.githubusercontent.com/herrbischoff/country-ip-blocks/master/ipv4/%s.txt", strings.ToLower(cc))
		resp, _ := http.Get(url)
		if resp != nil && resp.StatusCode == 200 {
			out, _ := os.Create(cachePath)
			bufio.NewReader(resp.Body).WriteTo(out)
			out.Close()
		}
	}
	
	var allIPs []string
	f, _ := os.Open(cachePath)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cidr := strings.TrimSpace(scanner.Text())
		if cidr != "" && !strings.HasPrefix(cidr, "#") {
			allIPs = append(allIPs, expandCIDR(cidr)...)
		}
	}
	return allIPs
}

func expandCIDR(cidr string) []string {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil { return nil }
	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incIP(ip) {
		ips = append(ips, ip.String())
	}
	return ips
}

func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 { break }
	}
}

func check(ip string, port int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), 500*time.Millisecond)
	if err != nil { return false }
	conn.Close()
	return true
}

func pickRandomIPs(pool []string, count int) []string {
	res := make([]string, count)
	for i := 0; i < count; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(pool))))
		res[i] = pool[n.Int64()]
	}
	return res
}
