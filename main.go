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
	"os/exec"
	"path/filepath"
	"runtime"
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
		fmt.Println("[+] poti v0.8.5 - Global Intelligence Radar Engine")
		fmt.Println("================================================================")
		fmt.Println(" [1] Pure Edge-to-Edge Global Random Scanner")
		fmt.Println(" [2] Target IP Intelligence Lookup Engine")
		fmt.Println(" [3] Wireless Wi-Fi Security & Gate Association")
		fmt.Println(" [4] Exit Terminal Framework")
		fmt.Println("================================================================")
		
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Select Operational Mode [1-4]: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			runScanner(reader)
		case "2":
			runLookup(reader)
		case "3":
			runWifiAudit(reader)
		case "4":
			fmt.Println("[*] Shutting down tactical asset grid pipeline.")
			os.Exit(0)
		default:
			fmt.Println("[-] Invalid matrix selection. Re-initializing shell layout...")
			time.Sleep(1 * time.Second)
		}
	}
}

func runWifiAudit(reader *bufio.Reader) {
	clearScreen()
	fmt.Println("[*] Entering Mode 3: Wireless Wi-Fi Security & Gate Association")
	fmt.Println("----------------------------------------------------------------")
	
	fmt.Println("[*] Phase 1: Locating native system binary infrastructure...")
	binaryPath := locateSystemWirelessTool()
	
	if binaryPath == "" {
		fmt.Println("[-] Critical Resolution Failure: No compliant wireless subsystem tools found.")
		fmt.Println("[-] Please ensure local hardware drivers or management utilities are configured.")
		fmt.Print("\nPress Enter to return to main menu...")
		reader.ReadString('\n')
		return
	}

	fmt.Printf("\n[+] Target Asset Located: %s\n", binaryPath)
	fmt.Println("[*] Phase 2: Scanning surrounding wireless environments...")
	time.Sleep(500 * time.Millisecond)
	
	scanSurroundingWifi(binaryPath)

	fmt.Print("\nEnter Target Wi-Fi SSID (Name) to connect: ")
	wifiName, _ := reader.ReadString('\n')
	wifiName = strings.TrimSpace(wifiName)

	if wifiName == "" {
		fmt.Println("[-] Operation aborted: SSID cannot be empty.")
		fmt.Print("\nPress Enter to return to main menu...")
		reader.ReadString('\n')
		return
	}

	fmt.Print("Enter Wi-Fi Authentication Password: ")
	wifiPass, _ := reader.ReadString('\n')
	wifiPass = strings.TrimSpace(wifiPass)

	clearScreen()
	fmt.Printf("[*] Initiating Hardware Association Protocol for SSID: [%s]\n", wifiName)
	fmt.Println("================================================================")
	
	success := connectToWifi(wifiName, wifiPass)
	if !success {
		fmt.Println("[-] Authentication Failure: Unable to associate with network.")
		fmt.Println("[-] Please verify credentials, security protocols, or range.")
		fmt.Print("\nPress Enter to return to main menu...")
		reader.ReadString('\n')
		return
	}

	fmt.Println("[+] Network Connection Established Successfully.")
	fmt.Println("[*] Gathering local network infrastructure configurations...")
	time.Sleep(1 * time.Second)

	gateways := fetchLocalGateways()
	var suggestedGateway string
	if len(gateways) > 0 {
		suggestedGateway = gateways[0]
	} else {
		suggestedGateway = "192.168.1.1"
	}

	client := &http.Client{Timeout: 4 * time.Second}
	resp, err := client.Get("https://ipinfo.io/json")
	var info IPInfo
	if err == nil && resp.StatusCode == 200 {
		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &info)
		resp.Body.Close()
	}

	riskLevel := "LOW RISK"
	vectorNotes := "Standard secure infrastructure network layer."
	if wifiPass == "" || strings.Contains(strings.ToLower(wifiName), "free") || strings.Contains(strings.ToLower(wifiName), "public") {
		riskLevel = "HIGH RISK"
		vectorNotes = "Unencrypted network environment. Cleartext traffic monitoring risk detected."
	}

	fmt.Println("\n================= WIRELESS NODE METADATA ======================")
	fmt.Printf("  Connected SSID  : %s\n", wifiName)
	fmt.Printf("  Local Gateway IP: %s\n", suggestedGateway)
	if info.IP != "" {
		fmt.Printf("  Egress Provider : %s\n", info.Org)
		fmt.Printf("  Assigned Region : %s, %s (%s)\n", info.City, info.Region, info.Country)
		fmt.Printf("  Geo Coordinates : %s\n", info.Loc)
	}
	fmt.Println("----------------------------------------------------------------")
	fmt.Printf("  SECURITY STATUS : [%s]\n", riskLevel)
	fmt.Printf("  Vulnerability   : %s\n", vectorNotes)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("[*] Management Access Gateways Discovered:")
	fmt.Printf("  --> HTTP Portal : http://%s\n", suggestedGateway)
	fmt.Printf("  --> HTTPS Portal: https://%s\n", suggestedGateway)
	fmt.Println("================================================================")
	fmt.Println("[*] Session verification complete.")
	
	fmt.Print("\nPress Enter to return to main menu...")
	reader.ReadString('\n')
}

func locateSystemWirelessTool() string {
	var targetBinary string
	var searchRoots []string

	switch runtime.GOOS {
	case "darwin":
		targetBinary = "airport"
		searchRoots = []string{
			"/System/Library/PrivateFrameworks/Apple80211.framework",
			"/usr/local/bin",
			"/usr/bin",
			"/opt",
		}
	case "linux":
		targetBinary = "nmcli"
		searchRoots = []string{
			"/usr/bin",
			"/bin",
			"/usr/sbin",
			"/sbin",
		}
	case "windows":
		targetBinary = "netsh.exe"
		systemRoot := os.Getenv("SystemRoot")
		if systemRoot == "" {
			systemRoot = "C:\\Windows"
		}
		searchRoots = []string{
			filepath.Join(systemRoot, "System32"),
			systemRoot,
		}
	default:
		return ""
	}

	if path, err := exec.LookPath(targetBinary); err == nil {
		return path
	}

	for _, root := range searchRoots {
		if _, err := os.Stat(root); os.IsNotExist(err) {
			continue
		}

		var foundPath string
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			
			if !info.IsDir() {
				dir := filepath.Dir(path)
				fmt.Printf("[*] Scanning system directory layout: %s\r", dir)
			}

			if !info.IsDir() && info.Name() == targetBinary {
				foundPath = path
				return filepath.SkipDir
			}
			return nil
		})

		if err == nil && foundPath != "" {
			fmt.Print("\n")
			return foundPath
		}
	}

	fmt.Print("\n")
	return ""
}

func connectToWifi(ssid, password string) bool {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("netsh", "wlan", "connect", "name="+ssid)
	case "linux":
		cmd = exec.Command("nmcli", "dev", "wifi", "connect", ssid, "password", password)
	case "darwin":
		cmd = exec.Command("networksetup", "-setairportnetwork", "en0", ssid, password)
	default:
		return false
	}
	err := cmd.Run()
	return err == nil
}

func scanSurroundingWifi(binaryPath string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command(binaryPath, "-s")
	case "linux":
		cmd = exec.Command(binaryPath, "-f", "SSID,BSSID,SECURITY,SIGNAL", "dev", "wifi")
	case "windows":
		cmd = exec.Command(binaryPath, "wlan", "show", "networks")
	default:
		fmt.Println("[-] Unsupported local operating system for hardware radio scanning.")
		return
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("[-] Hardware Interface Error: %v\n", err)
		fmt.Println("[-] Privilege restriction or interface disconnected. Unable to fetch real-time radio frames.")
		return
	}
	fmt.Println(string(output))
}

func fetchLocalGateways() []string {
	ifaces, err := net.Interfaces()
	if err != nil { return []string{} }
	
	var list []string
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 { continue }
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ipStr := ipnet.IP.String()
					if strings.HasPrefix(ipStr, "192.168.") || strings.HasPrefix(ipStr, "10.") {
						parts := strings.Split(ipStr, ".")
						if len(parts) == 4 {
							list = append(list, fmt.Sprintf("%s.%s.%s.1", parts[0], parts[1], parts[2]))
						}
					}
				}
			}
		}
	}
	return list
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
				if ip == "" { continue }

				wg.Add(1)
				semaphore <- struct{}{}

				go func(addr string) {
					defer wg.Done()
					defer func() { <-semaphore }()

					fmt.Printf("[...] Dialing: %s:80 -> establishing handshake...\n", addr)
					if check(addr, 80) {
						fmt.Printf("[!!!] Success: Active Node Verified -> %s\n", addr)
						select {
						case Scribble := <-results:
							_ = Scribble
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
		if err != nil { return "" }
		if b[0] == 0 || b[0] == 10 || b[0] == 127 { continue }
		if b[0] == 172 && (b[1] >= 16 && b[1] <= 31) { continue }
		if b[0] == 192 && b[1] == 168 { continue }
		if b[0] >= 224 { continue }
		return fmt.Sprintf("%d.%d.%d.%d", b[0], b[1], b[2], b[3])
	}
}

func check(ip string, port int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), 600*time.Millisecond)
	if err != nil { return false }
	conn.Close()
	return true
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
