package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/CYBERWOLF333/gopherrecon/pkg/banner"
	"github.com/CYBERWOLF333/gopherrecon/pkg/ctrecon"
	"github.com/CYBERWOLF333/gopherrecon/pkg/dnsrecon"
	"github.com/CYBERWOLF333/gopherrecon/pkg/headerrecon"
	"github.com/CYBERWOLF333/gopherrecon/pkg/reporter"
	"github.com/CYBERWOLF333/gopherrecon/pkg/tlsrecon"
)

const AppVersion = "1.0.0"

func runDemo() {
	banner.PrintBanner(AppVersion)
	fmt.Println("🚀 RUNNING BUILT-IN DEMO RECONNAISSANCE on target: 'cyberwolf-defense.io'")
	fmt.Println("   (Simulating live multi-vector passive OSINT scan...)")
	fmt.Println()

	// Artificial mini delay to make it feel authentic
	time.Sleep(300 * time.Millisecond)

	mockReport := &reporter.FullReconReport{
		TargetDomain: "cyberwolf-defense.io",
		ScanTime:     time.Now(),
		Duration:     "342ms",
		DNS: &dnsrecon.DNSReport{
			Domain:      "cyberwolf-defense.io",
			ARecords:    []string{"185.199.108.153", "185.199.109.153"},
			AAAARecords: []string{"2606:50c0:8000::153"},
			MXRecords:   []string{"aspmx.l.google.com (prio 1)", "alt1.aspmx.l.google.com (prio 5)"},
			NSRecords:   []string{"ns1.dnsimple.com", "ns2.dnsimple.com"},
			SPFRecord:   "v=spf1 include:_spf.google.com -all",
			SPFStatus:   "STRICT (-all: Hard fail unauthorized senders)",
			DMARCRecord: "v=DMARC1; p=reject; rua=mailto:dmarc@cyberwolf-defense.io",
			DMARCStatus: "STRONG (p=reject: Spoofed emails dropped)",
		},
		TLS: &tlsrecon.TLSReport{
			Active:        true,
			Version:       "TLS 1.3 (MODERN & RECOMMENDED)",
			CipherSuite:   "TLS_AES_128_GCM_SHA256",
			SubjectCN:     "cyberwolf-defense.io",
			IssuerOrg:     "Let's Encrypt",
			ValidFrom:     "2026-09-01",
			ValidUntil:    "2026-12-01",
			DaysRemaining: 51,
			IsExpired:     false,
			AltNames: []string{
				"cyberwolf-defense.io",
				"api.cyberwolf-defense.io",
				"portal.cyberwolf-defense.io",
				"vpn.cyberwolf-defense.io",
			},
		},
		Headers: &headerrecon.HeaderReport{
			URL:           "https://cyberwolf-defense.io",
			StatusCode:    200,
			SecurityScore: 90,
			Grade:         "A",
			Checks: []headerrecon.HeaderCheck{
				{HeaderName: "Strict-Transport-Security", Present: true, Value: "max-age=31536000; includeSubDomains; preload", Weight: 25, Severity: "HIGH"},
				{HeaderName: "Content-Security-Policy", Present: true, Value: "default-src 'self'; script-src 'self'", Weight: 25, Severity: "HIGH"},
				{HeaderName: "X-Frame-Options", Present: true, Value: "DENY", Weight: 15, Severity: "MEDIUM"},
				{HeaderName: "X-Content-Type-Options", Present: true, Value: "nosniff", Weight: 15, Severity: "MEDIUM"},
				{HeaderName: "Referrer-Policy", Present: true, Value: "strict-origin-when-cross-origin", Weight: 10, Severity: "LOW"},
				{HeaderName: "Permissions-Policy", Present: false, Weight: 10, Severity: "LOW", Remedy: "Add: Permissions-Policy: camera=(), microphone=()"},
			},
			ServerBanner: "cloudflare",
			InformationLeaks: []string{
				"Server: cloudflare",
			},
		},
		CTLogs: &ctrecon.CTReport{
			Domain:     "cyberwolf-defense.io",
			TotalFound: 5,
			Source:     "Certificate Transparency Logs (crt.sh)",
			SubdomainsDiscovered: []string{
				"admin.cyberwolf-defense.io",
				"api.cyberwolf-defense.io",
				"auth.cyberwolf-defense.io",
				"portal.cyberwolf-defense.io",
				"vpn.cyberwolf-defense.io",
			},
		},
	}

	reporter.PrintTerminalReport(mockReport)

	demoJson := "demo_recon_report.json"
	demoMd := "demo_recon_report.md"
	_ = reporter.ExportJSON(mockReport, demoJson)
	_ = reporter.ExportMarkdown(mockReport, demoMd)
	fmt.Printf("[+] Exported Demo artifacts: '%s' and '%s'\n", demoJson, demoMd)
}

func main() {
	domainFlag := flag.String("domain", "", "Target domain to recon (e.g. github.com)")
	demoFlag := flag.Bool("demo", false, "Run an interactive simulated recon audit")
	skipDNS := flag.Bool("skip-dns", false, "Skip DNS & Email security reconnaissance")
	skipTLS := flag.Bool("skip-tls", false, "Skip SSL/TLS certificate reconnaissance")
	skipHeaders := flag.Bool("skip-headers", false, "Skip HTTP security headers audit")
	skipCT := flag.Bool("skip-ct", false, "Skip Certificate Transparency logs enumeration")
	jsonPath := flag.String("json", "", "Export full report to JSON file")
	markdownPath := flag.String("markdown", "", "Export executive audit report to Markdown file")
	timeoutStr := flag.String("timeout", "5s", "Per-module network timeout")
	flag.Parse()

	if *demoFlag || (*domainFlag == "" && len(os.Args) == 1) {
		runDemo()
		return
	}

	if *domainFlag == "" {
		fmt.Println("Error: Target -domain is required! (Or use -demo)")
		flag.Usage()
		os.Exit(1)
	}

	target := strings.TrimSpace(strings.ToLower(*domainFlag))
	target = strings.TrimPrefix(target, "https://")
	target = strings.TrimPrefix(target, "http://")
	target = strings.TrimRight(target, "/")

	timeout, err := time.ParseDuration(*timeoutStr)
	if err != nil {
		fmt.Printf("Invalid timeout duration: %v\n", err)
		os.Exit(1)
	}

	banner.PrintBanner(AppVersion)
	fmt.Printf("[*] Initiating multi-threaded passive reconnaissance on: %s\n", banner.ColorCyan(target))

	startTime := time.Now()
	var wg sync.WaitGroup

	report := &reporter.FullReconReport{
		TargetDomain: target,
		ScanTime:     startTime,
	}

	// 1. DNS Recon Goroutine
	if !*skipDNS {
		wg.Add(1)
		go func() {
			defer wg.Done()
			report.DNS = dnsrecon.AnalyzeDNS(target, timeout)
		}()
	}

	// 2. TLS Recon Goroutine
	if !*skipTLS {
		wg.Add(1)
		go func() {
			defer wg.Done()
			report.TLS = tlsrecon.AnalyzeTLS(target, timeout)
		}()
	}

	// 3. Header Recon Goroutine
	if !*skipHeaders {
		wg.Add(1)
		go func() {
			defer wg.Done()
			report.Headers = headerrecon.AnalyzeHeaders(target, timeout)
		}()
	}

	// 4. CT Recon Goroutine
	if !*skipCT {
		wg.Add(1)
		go func() {
			defer wg.Done()
			report.CTLogs = ctrecon.DiscoverSubdomains(target, timeout)
		}()
	}

	// Wait for all 4 recon threads to finish concurrently
	wg.Wait()
	report.Duration = time.Since(startTime).Round(time.Millisecond).String()

	// Print Terminal Report
	reporter.PrintTerminalReport(report)

	// Export JSON if requested
	if *jsonPath != "" {
		if err := reporter.ExportJSON(report, *jsonPath); err != nil {
			fmt.Printf("[-] Failed to export JSON: %v\n", err)
		} else {
			fmt.Printf("[+] JSON report saved to: %s\n", *jsonPath)
		}
	}

	// Export Markdown if requested
	if *markdownPath != "" {
		if err := reporter.ExportMarkdown(report, *markdownPath); err != nil {
			fmt.Printf("[-] Failed to export Markdown: %v\n", err)
		} else {
			fmt.Printf("[+] Markdown audit report saved to: %s\n", *markdownPath)
		}
	}
}
