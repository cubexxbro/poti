package main

import (
	"bufio"
	"fmt"
	"net"
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

type ScanResult struct {
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
		fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: %s\r\nUser-Agent: poti/0.2.0\r\nConnection: close\r\n\r\n", ip)
	}

	scanner := bufio.NewScanner(conn)
	var bannerLines []string
	
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" && (port == 80 || port == 8080 || port == 443) {
			break
		}
		if len(bannerLines) < 8 {
			bannerLines = append(bannerLines, "    │ "+strings.TrimSpace(line))
		} else {
			break
		}
	}

	if len(bannerLines) > 0 {
		return strings.Join(bannerLines, "\n")
	}
	return "    │ [No immediate response / Silent service]"
}

func scanWorker(ip string, ports <-chan int, results chan<- ScanResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for port := range ports {
		banner := grabBanner(ip, port, 3*time.Second)
		if banner != "" {
			results <- ScanResult{Port: port, Banner: banner}
		}
	}
}

func main() {
	fmt.Printf("\033[1;36m%s\033[0m", asciiArt)
	fmt.Println("\033[1;32m[+] poti Engine - Active IP Monitoring System v0.2.0\033[0m")
	fmt.Println("--------------------------------------------------")

	targetIP := "127.0.0.1"
	portsToScan := []int{21, 22, 23, 25, 53, 80, 110, 143, 443, 8080, 8888}

	fmt.Printf("[*] Target IP: %s\n", targetIP)
	fmt.Printf("[*] Scanning & Grabbing banners from %d ports...\n\n", len(portsToScan))

	numWorkers := 5
	portsChan := make(chan int, len(portsToScan))
	resultsChan := make(chan ScanResult, len(portsToScan))

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go scanWorker(targetIP, portsChan, resultsChan, &wg)
	}

	for _, port := range portsToScan {
		portsChan <- port
	}
	close(portsChan)

	wg.Wait()
	close(resultsChan)

	fmt.Println("----------------- Discovery Report -----------------")
	found := 0
	for result := range resultsChan {
		found++
		fmt.Printf(" [🔥] Port \033[1;33m%d\033[0m is OPEN\n", result.Port)
		fmt.Println(result.Banner)
		fmt.Println()
	}

	if found == 0 {
		fmt.Println(" [-] No active services responded in this cycle.")
	}
	fmt.Println("--------------------------------------------------")
	fmt.Println("\n[*] Monitoring cycle finished. Awaiting next pulse...")
}
