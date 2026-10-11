package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

// Define our Intel Report Struct with JSON tags
type ThreatIntelReport struct {
	TargetDomain   string    `json:"target_domain"`
	RiskScore      int       `json:"risk_score"`
	Vulnerabilities []string `json:"vulnerabilities"`
	ReportGenerated time.Time `json:"report_generated"`
}

func main() {
	fmt.Println("=== ✉️ Lesson 10: Files, JSON & HTTP Networking ===")

	// 1. Creating and Packing JSON (Marshal)
	report := ThreatIntelReport{
		TargetDomain:   "cyber-lab.corp",
		RiskScore:      42,
		Vulnerabilities: []string{"Outdated TLS 1.0", "Missing HSTS Header", "Exposed Debug Port"},
		ReportGenerated: time.Now(),
	}

	jsonData, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		fmt.Println("Failed to pack JSON:", err)
		return
	}
	fmt.Println("--- 📜 Serialized JSON Report ---")
	fmt.Println(string(jsonData))

	// 2. Writing to a local file and reading it back
	reportFileName := "sample_report.json"
	err = os.WriteFile(reportFileName, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing report file:", err)
		return
	}
	defer os.Remove(reportFileName) // Clean up file after run
	fmt.Printf("Saved report to '%s' successfully!\n", reportFileName)

	// Unpacking JSON back into a Go struct (Unmarshal)
	readBytes, err := os.ReadFile(reportFileName)
	if err != nil {
		fmt.Println("Error reading back file:", err)
		return
	}

	var loadedReport ThreatIntelReport
	err = json.Unmarshal(readBytes, &loadedReport)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	fmt.Printf("Unpacked back from disk! Target: %s, Vulns Found: %d\n",
		loadedReport.TargetDomain, len(loadedReport.Vulnerabilities))

	// 3. HTTP Server and Client in Action!
	fmt.Println("\n--- 🌐 HTTP Server & Client Mock Handshake ---")
	// Using httptest to spin up a safe, lightning-fast in-memory web server!
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Security-Header", "GopherArmor-v1")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"online","defense_mode":"active"}`))
	}))
	defer testServer.Close()

	fmt.Printf("Test HTTP server started at: %s\n", testServer.URL)

	// Make an HTTP GET request to our server
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(testServer.URL)
	if err != nil {
		fmt.Println("Failed to GET:", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Printf("Response Status Code: %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
	fmt.Printf("Custom Security Header: %s\n", resp.Header.Get("X-Security-Header"))
	fmt.Printf("Response Body: %s\n", string(bodyBytes))
}
