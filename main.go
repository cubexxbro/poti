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
	fmt.Println(`
  _____   ____   _______  _____ 
 |  __ \ / __ \ |__   __||_   _|
 | |__) | |  | |   | |     | |  
 |  ___/| |  | |   | |     | |  
 | |    | |__| |   | |    _| |_ 
 |_|     \____/    |_|   |_____|
                                `)
	fmt.Println("[+] poti v0.8.0 - Global Intelligence Radar")
	fmt.Println("[*] Mode: Verbose Global Scanning")
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Target CC (e.g. US, CN, JP or ALL): ")
	cc, _ := reader.ReadString('\n')
	cc = strings.TrimSpace(strings.ToUpper(cc))
	if cc == "" { cc = "ALL" }

	fmt.Println("[*] Initializing global network asset lattice...")
	ips := loadGlobalAssets(cc)
	if len(ips) == 0 { fmt.Println("[-] No assets fetched."); return }

	fmt.Print("Quota: ")
	q, _ := reader.ReadString('\n')
	limit := 100
	fmt.Sscanf(strings.TrimSpace(q), "%d", &limit)

	fmt.Println("[*] Execution logic loaded. Terminal pipeline active.")
	results := make(chan string, limit)
	var wg sync.WaitGroup
	
	semaphore := make(chan struct{}, 20) 

	for _, ip := range pickRandomIPs(ips, limit) {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(addr string) {
			defer wg.Done()
			defer func() { <-semaphore }()
			
			fmt.Printf("[...] Dialing: %s:80 -> establishing handshake...\n", addr)
			if check(addr, 80) {
				fmt.Printf("[!!!] Success: Active Node Verified -> %s\n", addr)
				results <- addr
			} else {
				fmt.Printf("[-] Failed: Connection dropped for %s\n", addr)
			}
			time.Sleep(150 * time.Millisecond) 
		}(ip)
	}

	go func() { wg.Wait(); close(results) }()
	
	var activeCount int
	for range results { activeCount++ }
	fmt.Printf("[*] Radar operation finished. Total active assets found: %d\n", activeCount)
}

func loadGlobalAssets(cc string) []string {
	var ccs []string
	if cc == "ALL" {
		ccs = []string{"us", "cn", "jp", "de", "kr", "gb", "fr", "ru"}
	} else {
		ccs = []string{strings.ToLower(cc)}
	}

	var allIPs []string
	for _, currentCC := range ccs {
		cache := filepath.Join(os.TempDir(), "poti_"+currentCC+".cache")
		if _, err := os.Stat(cache); os.IsNotExist(err) {
			fmt.Printf("[*] Downloading remote registry files for sovereign space: %s...\n", strings.ToUpper(currentCC))
			url := fmt.Sprintf("https://raw.githubusercontent.com/herrbischoff/country-ip-blocks/master/ipv4/%s.txt", currentCC)
			resp, err := http.Get(url)
			if err != nil || resp.StatusCode != 200 {
				fmt.Printf("[-] Skip: Failed to sync database for registry zone %s\n", strings.ToUpper(currentCC))
				continue
			}
			out, _ := os.Create(cache)
			bufio.NewReader(resp.Body).WriteTo(out)
			out.Close()
		}

		f, err := os.Open(cache)
		if err != nil { continue }
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" && !strings.HasPrefix(line, "#") {
				ip, ipnet, err := net.ParseCIDR(line)
				if err == nil {
					for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incIP(ip) {
						allIPs = append(allIPs, ip.String())
					}
				}
			}
		}
		f.Close()
	}
	return allIPs
}

func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 { break }
	}
}

func check(ip string, port int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), 600*time.Millisecond)
	if err != nil { return false }
	conn.Close()
	return true
}

func pickRandomIPs(pool []string, count int) []string {
	if len(pool) == 0 { return nil }
	if count > len(pool) { count = len(pool) }
	res := make([]string, count)
	for i := 0; i < count; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(pool))))
		res[i] = pool[n.Int64()]
	}
	return res
}
