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

func main() {
	fmt.Println("[+] poti v0.7.4 - Turbo Mode")
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Target CC: ")
	cc, _ := reader.ReadString('\n')
	cc = strings.TrimSpace(strings.ToUpper(cc))
	if cc == "" { cc = "CN" }

	ips := getIPs(cc)
	if len(ips) == 0 { fmt.Println("[-] Fail."); return }

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

func getIPs(cc string) []string {
	return []string{"1.1.1.1", "8.8.8.8"}
}

func check(ip string, port int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), 800*time.Millisecond)
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
