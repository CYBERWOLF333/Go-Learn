package ctrecon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"
)

type CrtShEntry struct {
	IssuerCaID     int    `json:"issuer_ca_id"`
	IssuerName     string `json:"issuer_name"`
	CommonName     string `json:"common_name"`
	NameValue      string `json:"name_value"`
	ID             int64  `json:"id"`
	EntryTimestamp string `json:"entry_timestamp"`
	NotBefore      string `json:"not_before"`
	NotAfter       string `json:"not_after"`
}

type CTReport struct {
	Domain              string   `json:"domain"`
	SubdomainsDiscovered []string `json:"subdomains_discovered"`
	TotalFound          int      `json:"total_found"`
	Source              string   `json:"source"`
	Error               string   `json:"error,omitempty"`
}

// DiscoverSubdomains queries public Certificate Transparency logs passively
func DiscoverSubdomains(domain string, timeout time.Duration) *CTReport {
	report := &CTReport{
		Domain: domain,
		Source: "Certificate Transparency Logs (crt.sh)",
	}

	url := fmt.Sprintf("https://crt.sh/?q=%%25.%s&output=json", domain)

	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		report.Error = fmt.Sprintf("Failed to build CT request: %v", err)
		return report
	}
	req.Header.Set("User-Agent", "GopherRecon-OSINT-Parser/1.0")

	resp, err := client.Do(req)
	if err != nil {
		report.Error = fmt.Sprintf("CT Logs unavailable or restricted: %v", err)
		return report
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		report.Error = fmt.Sprintf("crt.sh returned HTTP status %d", resp.StatusCode)
		return report
	}

	var entries []CrtShEntry
	if err := json.NewDecoder(resp.Body).Decode(&entries); err != nil {
		report.Error = fmt.Sprintf("Failed to parse crt.sh JSON: %v", err)
		return report
	}

	subdomainSet := make(map[string]struct{})
	for _, entry := range entries {
		// crt.sh can return multiple names separated by newlines
		names := strings.Split(entry.NameValue, "\n")
		for _, name := range names {
			clean := strings.ToLower(strings.TrimSpace(name))
			clean = strings.TrimPrefix(clean, "*.")
			if strings.HasSuffix(clean, domain) && clean != domain && !strings.Contains(clean, "@") {
				subdomainSet[clean] = struct{}{}
			}
		}
	}

	var list []string
	for sub := range subdomainSet {
		list = append(list, sub)
	}
	sort.Strings(list)

	report.SubdomainsDiscovered = list
	report.TotalFound = len(list)
	return report
}
