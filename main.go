package main

import (
	"fmt"
	"net"
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

func scanPort(ip string, port int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return
	}
	conn.Close()

	results <- port
}

func main() {
	fmt.Printf("\033[1;36m%s\033[0m", asciiArt)
	fmt.Println("\033[1;32m[+] poti Engine - Active IP Monitoring System v0.1.0\033[0m")
	fmt.Println("--------------------------------------------------")

	targetIP := "127.0.0.1"
	portsToScan := []int{21, 22, 80, 443, 8080, 8888}

	fmt.Printf("[*] Target IP: %s\n", targetIP)
	fmt.Printf("[*] Scanning %d ports concurrently...\n\n", len(portsToScan))

	var wg sync.WaitGroup
	results := make(chan int, len(portsToScan))

	for _, port := range portsToScan {
		wg.Add(1)
		go scanPort(targetIP, port, &wg, results)
	}

	wg.Wait()
	close(results)

	fmt.Println("----------------- Scan Results -----------------")
	found := 0
	for openPort := range results {
		found++
		fmt.Printf(" [🔥] Port \033[1;33m%d\033[0m is OPEN on %s\n", openPort, targetIP)
	}

	if found == 0 {
		fmt.Println(" [-] No open ports discovered in this cycle.")
	}
	fmt.Println("--------------------------------------------------")
	fmt.Println("\n[*] Monitoring cycle finished. Awaiting next pulse...")
}
