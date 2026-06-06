package main

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type IPInfo struct {
	IP           string `json:"ip"`
	City         string `json:"city"`
	Region       string `json:"region"`
	Country      string `json:"country"`
	Loc          string `json:"loc"`
	Org          string `json:"org"`
	Postal       string `json:"postal"`
	Timezone     string `json:"timezone"`
}

func main() {
	for {
		clearScreen()
		fmt.Println(`
  _____   ____   _______  _____ 
 |  __ \ / __ \ |__   __||_   _|
 | |__) | |  | |   | |     | |  
 |  ___/| |  | |   | |     | |  
 | |    | |__| |   | |    _| |_ 
 |_|     \____/    |_|   |_____|
                                `)
		fmt.Println("[+] poti v0.8.1 - Global Intelligence Radar Engine")
		fmt.Println("==================================================")
		fmt.Println(" [1] Pure Edge-to-Edge Global Random Scanner")
		fmt.Println(" [2] Target IP Intelligence Lookup Engine")
		fmt.Println(" [3] Exit Terminal Framework")
		fmt.Println("==================================================")
		
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Select Operational Mode [1-3]: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			runScanner(reader)
		case "2":
			runLookup(reader)
		case "3":
			fmt.Println("[*] Shutting down tactical asset grid pipeline.")
			os.Exit(0)
		default:
			fmt.Println("[-] Invalid matrix selection. Re-initializing shell layout...")
			time.Sleep(1 * time.Second)
		}
	}
}

func runScanner(reader *bufio.Reader) {
	clearScreen()
	fmt.Println("[*] Entering Mode 1: Global Edge Scanner")
	fmt.Println("--------------------------------------------------")
	fmt.Print("Quota (Total targeted active nodes to find before auto-exit): ")
	q, _ := reader.ReadString('\n')
	limit := 100
	fmt.Sscanf(strings.TrimSpace(q), "%d", &limit)

	fmt.Println("[*] Activating mathematical global distribution grid...")
	fmt.Println("[*] Execution logic loaded. Terminal pipeline active.")

	results := make(chan string)
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 50) 
	stopSignal := make(chan struct{})

	go func() {
		for {
			select {
			case <-stopSignal:
				return
			default:
				ip := generateGlobalIP()
				if ip == "" {
					continue
				}

				wg.Add(1)
				semaphore <- struct{}{}

				go func(addr string) {
					defer wg.Done()
					defer func() { <-semaphore }()

					fmt.Printf("[...] Dialing: %s:80 -> establishing handshake...\n", addr)
					if check(addr, 80) {
						fmt.Printf("[!!!] Success: Active Node Verified -> %s\n", addr)
						select {
						case results <- addr:
						case <-stopSignal:
						}
					} else {
						fmt.Printf("[-] Failed: Connection dropped for %s\n", addr)
					}
					time.Sleep(100 * time.Millisecond)
				}(ip)
			}
		}
	}()

	var activeCount int
	for range results {
		activeCount++
		if activeCount >= limit {
			close(stopSignal)
			break
		}
	}

	wg.Wait()
	close(results)

	fmt.Printf("\n[*] Radar operation finished. Total active assets verified: %d\n", activeCount)
	fmt.Print("Press Enter to return to main menu...")
	reader.ReadString('\n')
}

func runLookup(reader *bufio.Reader) {
	clearScreen()
	fmt.Println("[*] Entering Mode 2: IP Intelligence Lookup")
	fmt.Println("--------------------------------------------------")
	fmt.Print("Enter Target IP Address: ")
	ipInput, _ := reader.ReadString('\n')
	ipInput = strings.TrimSpace(ipInput)

	if net.ParseIP(ipInput) == nil {
		fmt.Println("[-] Absolute parsing failure: Token string is not a valid IPv4/IPv6 layout.")
		fmt.Print("\nPress Enter to return to main menu...")
		reader.ReadString('\n')
		return
	}

	fmt.Printf("[*] Querying edge database registry paths for payload: %s...\n", ipInput)
	
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(fmt.Sprintf("https://ipinfo.io/%s/json", ipInput))
	if err != nil {
		fmt.Printf("[-] API Link Interrupted: Unable to fetch geolocation structures -> %v\n", err)
		fmt.Print("\nPress Enter to return to main menu...")
		reader.ReadString('\n')
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("[-] Server returned operational error metric code: %d\n", resp.StatusCode)
		fmt.Print("\nPress Enter to return to main menu...")
		reader.ReadString('\n')
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[-] Infrastructure payload stream parsing crash.")
		fmt.Print("\nPress Enter to return to main menu...")
		reader.ReadString('\n')
		return
	}

	var info IPInfo
	if err := json.Unmarshal(body, &info); err != nil {
		fmt.Println("[-] Metadata deserialization structurally corrupted.")
		fmt.Print("\nPress Enter to return to main menu...")
		reader.ReadString('\n')
		return
	}

	fmt.Println("\n================ TARGET INTELLIGENCE METADATA ================")
	fmt.Printf("  IP Address  : %s\n", info.IP)
	fmt.Printf("  Country     : %s\n", info.Country)
	fmt.Printf("  Region/State: %s\n", info.Region)
	fmt.Printf("  City/Zone   : %s\n", info.City)
	fmt.Printf("  Coordinates : %s\n", info.Loc)
	fmt.Printf("  ISP / Org   : %s\n", info.Org)
	fmt.Printf("  Postal Code : %s\n", info.Postal)
	fmt.Printf("  Timezone    : %s\n", info.Timezone)
	fmt.Println("==============================================================")
	fmt.Println("[*] Data retrieval cycle executed successfully.\n")
	
	fmt.Print("Press Enter to return to main menu...")
	reader.ReadString('\n')
}

func generateGlobalIP() string {
	b := make([]byte, 4)
	for {
		_, err := rand.Read(b)
		if err != nil {
			return ""
		}

		if b[0] == 0 || b[0] == 10 || b[0] == 127 {
			continue
		}
		if b[0] == 172 && (b[1] >= 16 && b[1] <= 31) {
			continue
		}
		if b[0] == 192 && b[1] == 168 {
			continue
		}
		if b[0] >= 224 {
			continue
		}

		return fmt.Sprintf("%d.%d.%d.%d", b[0], b[1], b[2], b[3])
	}
}

func check(ip string, port int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), 600*time.Millisecond)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
