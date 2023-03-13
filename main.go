package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// Set up command line arguments
var (
	domain  string
	output  string
	baseURL string = "https://crt.sh/?q=%s&output=json"
)

func init() {
	flag.StringVar(&domain, "d", "", "Domain to search for domains")
	flag.StringVar(&output, "o", "", "Output file to write domains to")
	flag.Parse()
}

// Get domains from crt.sh
func getDomains() ([]map[string]interface{}, error) {
	resp, err := http.Get(fmt.Sprintf(baseURL, domain))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var domains []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&domains)
	if err != nil {
		return nil, err
	}

	return domains, nil
}

// Filter domains to remove duplicates and unwanted domains
func filterDomains(domains []map[string]interface{}) []string {
	var domainCollection []string
	filters := []string{"*", "\\", "-", "@"}

	for _, d := range domains {
		domainName := strings.TrimSpace(d["name_value"].(string))
		domainSplit := strings.Split(domainName, "\n")

		if len(domainSplit) == 1 && !containsAny(domainName, filters) {
			domainCollection = append(domainCollection, domainName)
		} else {
			for _, d2 := range domainSplit {
				if !containsAny(d2, filters) {
					domainCollection = append(domainCollection, d2)
				}
			}
		}
	}

	return uniqueStrings(domainCollection)
}

// Write domains to file
func writeDomains(domains []string) error {
	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, domain := range domains {
		file.WriteString(domain + "\n")
	}

	return nil
}

// Check if a string contains any of the substrings in a list
func containsAny(str string, substrings []string) bool {
	for _, s := range substrings {
		if strings.Contains(str, s) {
			return true
		}
	}
	return false
}

// Remove duplicates from a string slice
func uniqueStrings(strings []string) []string {
	encountered := map[string]bool{}
	result := []string{}
	for _, s := range strings {
		if !encountered[s] {
			encountered[s] = true
			result = append(result, s)
		}
	}
	return result
}

// Main function
func main() {
	if domain == "" {
		log.Fatalln("Domain is required.")
	}

	domains, err := getDomains()
	if err != nil {
		log.Fatalln("Error getting domains from crt.sh:", err)
	}

	domainCollection := filterDomains(domains)

	for _, d := range domainCollection {
		fmt.Println(d)
	}

	if output != "" {
		err = writeDomains(domainCollection)
		if err != nil {
			log.Fatalln("Error writing domains to file:", err)
		}
	}
}
