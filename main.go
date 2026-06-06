package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
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

type ShodanHost struct {
	IP        string   `json:"ip_str"`
	Port      int      `json:"port"`
	Org       string   `json:"org"`
	Data      string   `json:"data"`
	Transport string   `json:"transport"`
}

type ShodanResponse struct {
	Matches []ShodanHost `json:"matches"`
	Total   int          `json:"total"`
}

func searchShodan(query string, apiKey string) (*ShodanResponse, error) {
	apiURL := fmt.Sprintf("https://api.shodan.io/shodan/host/search?key=%s&query=%s", apiKey, url.QueryEscape(query))
	
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	var result ShodanResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func main() {
	fmt.Printf("\033[1;36m%s\033[0m", asciiArt)
	fmt.Println("\033[1;32m[+] poti Engine - Space Mapping Search v0.3.0\033[0m")
	fmt.Println("--------------------------------------------------")

	queryFlag := flag.String("q", "", "Search query (keyword, port, or IP)")
	apiKeyFlag := flag.String("key", "", "Shodan API Key")
	flag.Parse()

	apiKey := *apiKeyFlag
	if apiKey == "" {
		apiKey = os.Getenv("SHODAN_API_KEY")
	}

	if *queryFlag == "" {
		fmt.Println("\033[1;31m[-] Error: Search query (-q) is required.\033[0m")
		fmt.Println("Usage:")
		fmt.Println("  ./poti -q \"product:Apache\" -key \"YOUR_API_KEY\"")
		fmt.Println("  ./poti -q \"8.8.8.8\"")
		fmt.Println("\n*Or set environment variable: export SHODAN_API_KEY=your_key")
		fmt.Println("--------------------------------------------------")
		return
	}

	if apiKey == "" {
		fmt.Println("\033[1;31m[-] Error: Shodan API key is missing.\033[0m")
		fmt.Println("Please provide it via -key or SHODAN_API_KEY environment variable.")
		fmt.Println("--------------------------------------------------")
		return
	}

	fmt.Printf("[*] Querying global intelligence for: \"%s\"...\n\n", *queryFlag)

	results, err := searchShodan(*queryFlag, apiKey)
	if err != nil {
		fmt.Printf("\033[1;31m[-] API Request Failed: %v\033[0m\n", err)
		return
	}

	fmt.Printf("----------------- Discovery Report (Total: %d) -----------------\n", results.Total)
	
	for _, match := range results.Matches {
		fmt.Printf(" [🔥] Target: \033[1;33m%s\033[0m:\033[1;32m%d\033[0m (%s/%s)\n", match.IP, match.Port, match.Transport, match.Org)
		
		lines := strings.Split(strings.TrimSpace(match.Data), "\n")
		for i, line := range lines {
			if i < 5 {
				fmt.Printf("    │ %s\n", strings.TrimSpace(line))
			} else {
				fmt.Println("    │ ... [Truncated]")
				break
			}
		}
		fmt.Println()
	}
	fmt.Println("--------------------------------------------------")
}
