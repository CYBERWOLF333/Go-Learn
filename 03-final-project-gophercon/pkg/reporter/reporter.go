package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/CYBERWOLF333/gopherrecon/pkg/banner"
	"github.com/CYBERWOLF333/gopherrecon/pkg/ctrecon"
	"github.com/CYBERWOLF333/gopherrecon/pkg/dnsrecon"
	"github.com/CYBERWOLF333/gopherrecon/pkg/headerrecon"
	"github.com/CYBERWOLF333/gopherrecon/pkg/tlsrecon"
)

type FullReconReport struct {
	TargetDomain string                    `json:"target_domain"`
	ScanTime     time.Time                 `json:"scan_time"`
	Duration     string                    `json:"duration"`
	DNS          *dnsrecon.DNSReport       `json:"dns,omitempty"`
	TLS          *tlsrecon.TLSReport       `json:"tls,omitempty"`
	Headers      *headerrecon.HeaderReport `json:"headers,omitempty"`
	CTLogs       *ctrecon.CTReport         `json:"certificate_transparency,omitempty"`
}

// PrintTerminalReport prints a beautiful ANSI colored executive summary
func PrintTerminalReport(r *FullReconReport) {
	fmt.Println(banner.ColorBold("\n================================================================="))
	fmt.Printf(" 🎯 RECONNAISSANCE TARGET: %s\n", banner.ColorCyan(r.TargetDomain))
	fmt.Printf(" 🕒 Scan Completed: %s (Duration: %s)\n", r.ScanTime.Format("2006-01-02 15:04:05 MST"), r.Duration)
	fmt.Println(banner.ColorBold("================================================================="))

	// 1. DNS Section
	if r.DNS != nil {
		fmt.Printf("\n%s\n", banner.ColorBold("🌐 [1] DNS INTEL & EMAIL SECURITY"))
		fmt.Println("-----------------------------------------------------------------")
		if len(r.DNS.ARecords) > 0 {
			fmt.Printf("  IPv4 (A) Records:     %s\n", strings.Join(r.DNS.ARecords, ", "))
		}
		if len(r.DNS.AAAARecords) > 0 {
			fmt.Printf("  IPv6 (AAAA) Records:  %s\n", strings.Join(r.DNS.AAAARecords, ", "))
		}
		if len(r.DNS.MXRecords) > 0 {
			fmt.Printf("  Mail (MX) Servers:    %s\n", strings.Join(r.DNS.MXRecords, " | "))
		}
		if len(r.DNS.NSRecords) > 0 {
			fmt.Printf("  Name Servers (NS):    %s\n", strings.Join(r.DNS.NSRecords, ", "))
		}

		// Email Security posture
		spfColor := banner.ColorGreen
		if strings.Contains(r.DNS.SPFStatus, "MISSING") || strings.Contains(r.DNS.SPFStatus, "WEAK") {
			spfColor = banner.ColorRed
		}
		fmt.Printf("  SPF Policy Status:    %s\n", spfColor(r.DNS.SPFStatus))

		dmarcColor := banner.ColorGreen
		if strings.Contains(r.DNS.DMARCStatus, "MISSING") || strings.Contains(r.DNS.DMARCStatus, "MONITORING") {
			dmarcColor = banner.ColorYellow
		}
		fmt.Printf("  DMARC Status:         %s\n", dmarcColor(r.DNS.DMARCStatus))
	}

	// 2. TLS Section
	if r.TLS != nil {
		fmt.Printf("\n%s\n", banner.ColorBold("🔒 [2] SSL/TLS CRYPTOGRAPHIC POSTURE"))
		fmt.Println("-----------------------------------------------------------------")
		if !r.TLS.Active {
			fmt.Printf("  Status: %s\n", banner.ColorRed("INACTIVE / OFFLINE ("+r.TLS.Error+")"))
		} else {
			fmt.Printf("  Protocol Version:     %s\n", banner.ColorCyan(r.TLS.Version))
			fmt.Printf("  Cipher Suite:         %s\n", r.TLS.CipherSuite)
			fmt.Printf("  Certificate Subject:  %s\n", r.TLS.SubjectCN)
			fmt.Printf("  Certificate Issuer:   %s\n", r.TLS.IssuerOrg)
			fmt.Printf("  Validity Window:      %s to %s\n", r.TLS.ValidFrom, r.TLS.ValidUntil)

			expiryColor := banner.ColorGreen
			if r.TLS.DaysRemaining < 30 {
				expiryColor = banner.ColorRed
			}
			fmt.Printf("  Days Until Expiry:    %s days\n", expiryColor(fmt.Sprintf("%d", r.TLS.DaysRemaining)))

			if len(r.TLS.AltNames) > 0 {
				fmt.Printf("  SAN Subdomains:       %s\n", banner.ColorYellow(strings.Join(r.TLS.AltNames, ", ")))
			}
		}
	}

	// 3. HTTP Security Headers Section
	if r.Headers != nil {
		fmt.Printf("\n%s\n", banner.ColorBold("🛡️  [3] HTTP DEFENSIVE SECURITY HEADERS"))
		fmt.Println("-----------------------------------------------------------------")
		if r.Headers.Error != "" {
			fmt.Printf("  Status: %s\n", banner.ColorRed("Error: "+r.Headers.Error))
		} else {
			gradeColor := banner.ColorGreen
			if r.Headers.Grade == "C" || r.Headers.Grade == "D" {
				gradeColor = banner.ColorYellow
			} else if r.Headers.Grade == "F" {
				gradeColor = banner.ColorRed
			}

			fmt.Printf("  Security Score:       %d / 100\n", r.Headers.SecurityScore)
			fmt.Printf("  Audit Grade:          [%s]\n", gradeColor(banner.ColorBold(r.Headers.Grade)))
			fmt.Printf("  HTTP Response Status: %d\n", r.Headers.StatusCode)

			for _, c := range r.Headers.Checks {
				if c.Present {
					fmt.Printf("  %s %-26s : %s\n", banner.ColorGreen("✓"), c.HeaderName, banner.Dim+c.Value+banner.Reset)
				} else {
					fmt.Printf("  %s %-26s : %s\n",
						banner.ColorRed("✗"), c.HeaderName, banner.ColorYellow(fmt.Sprintf("MISSING (-%d pts, %s)", c.Weight, c.Severity)))
				}
			}

			if len(r.Headers.InformationLeaks) > 0 {
				fmt.Println("\n  ⚠️  Information Leaks (Server Banners):")
				for _, leak := range r.Headers.InformationLeaks {
					fmt.Printf("     ! %s\n", banner.ColorYellow(leak))
				}
			}
		}
	}

	// 4. CT Logs Section
	if r.CTLogs != nil {
		fmt.Printf("\n%s\n", banner.ColorBold("🛰️  [4] PASSIVE CERTIFICATE TRANSPARENCY ENUMERATION"))
		fmt.Println("-----------------------------------------------------------------")
		if r.CTLogs.Error != "" {
			fmt.Printf("  Status: %s\n", banner.Dim+"("+r.CTLogs.Error+")"+banner.Reset)
		} else {
			fmt.Printf("  Total Subdomains Discovered: %s\n", banner.ColorGreen(fmt.Sprintf("%d", r.CTLogs.TotalFound)))
			displayLimit := 10
			for i, sub := range r.CTLogs.SubdomainsDiscovered {
				if i >= displayLimit {
					fmt.Printf("  ... and %d more subdomains (see JSON or Markdown export!)\n",
						len(r.CTLogs.SubdomainsDiscovered)-displayLimit)
					break
				}
				fmt.Printf("  -> %s\n", sub)
			}
		}
	}

	fmt.Println(banner.ColorBold("\n================================================================="))
	fmt.Println(" ✨ Reconnaissance completed successfully! Stay ethical.")
	fmt.Println(banner.ColorBold("=================================================================\n"))
}

// ExportJSON writes the report to a JSON file
func ExportJSON(report *FullReconReport, filePath string) error {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

// ExportMarkdown writes the report as a clean Markdown audit document
func ExportMarkdown(r *FullReconReport, filePath string) error {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# 🛡️ OSINT & Security Reconnaissance Report: `%s`\n\n", r.TargetDomain))
	sb.WriteString(fmt.Sprintf("- **Scan Timestamp:** %s\n", r.ScanTime.Format("2006-01-02 15:04:05 MST")))
	sb.WriteString(fmt.Sprintf("- **Scan Duration:** %s\n", r.Duration))
	sb.WriteString("- **Assessment Engine:** [GopherRecon](https://github.com/CYBERWOLF333/gopherrecon)\n\n")
	sb.WriteString("---\n\n")

	// DNS Section
	if r.DNS != nil {
		sb.WriteString("## 1. 🌐 DNS & Email Security\n\n")
		sb.WriteString("| Record Type | Values |\n")
		sb.WriteString("| :--- | :--- |\n")
		sb.WriteString(fmt.Sprintf("| **IPv4 (A)** | `%s` |\n", strings.Join(r.DNS.ARecords, "`, `")))
		sb.WriteString(fmt.Sprintf("| **IPv6 (AAAA)** | `%s` |\n", strings.Join(r.DNS.AAAARecords, "`, `")))
		sb.WriteString(fmt.Sprintf("| **Mail (MX)** | `%s` |\n", strings.Join(r.DNS.MXRecords, "<br>")))
		sb.WriteString(fmt.Sprintf("| **Nameservers (NS)** | `%s` |\n", strings.Join(r.DNS.NSRecords, "`, `")))
		sb.WriteString(fmt.Sprintf("| **SPF Policy** | %s (`%s`) |\n", r.DNS.SPFStatus, r.DNS.SPFRecord))
		sb.WriteString(fmt.Sprintf("| **DMARC Policy** | %s (`%s`) |\n\n", r.DNS.DMARCStatus, r.DNS.DMARCRecord))
	}

	// TLS Section
	if r.TLS != nil && r.TLS.Active {
		sb.WriteString("## 2. 🔒 SSL/TLS Cryptographic Posture\n\n")
		sb.WriteString(fmt.Sprintf("- **Version:** `%s`\n", r.TLS.Version))
		sb.WriteString(fmt.Sprintf("- **Cipher Suite:** `%s`\n", r.TLS.CipherSuite))
		sb.WriteString(fmt.Sprintf("- **Subject Common Name:** `%s`\n", r.TLS.SubjectCN))
		sb.WriteString(fmt.Sprintf("- **Issuer:** `%s`\n", r.TLS.IssuerOrg))
		sb.WriteString(fmt.Sprintf("- **Validity:** %s to %s (%d days remaining)\n\n", r.TLS.ValidFrom, r.TLS.ValidUntil, r.TLS.DaysRemaining))
		if len(r.TLS.AltNames) > 0 {
			sb.WriteString(fmt.Sprintf("- **Subject Alternative Names (SANs):** `%s`\n\n", strings.Join(r.TLS.AltNames, "`, `")))
		}
	}

	// HTTP Security Headers
	if r.Headers != nil && r.Headers.Error == "" {
		sb.WriteString("## 3. 🛡️ HTTP Security Headers Audit\n\n")
		sb.WriteString(fmt.Sprintf("### Security Score: **%d / 100** | Grade: **%s**\n\n", r.Headers.SecurityScore, r.Headers.Grade))
		sb.WriteString("| Header | Status | Severity | Details / Remedy |\n")
		sb.WriteString("| :--- | :---: | :---: | :--- |\n")
		for _, c := range r.Headers.Checks {
			if c.Present {
				sb.WriteString(fmt.Sprintf("| `%s` | ✅ Present | - | `%s` |\n", c.HeaderName, c.Value))
			} else {
				sb.WriteString(fmt.Sprintf("| `%s` | ❌ Missing | %s | %s |\n", c.HeaderName, c.Severity, c.Remedy))
			}
		}
		sb.WriteString("\n")
	}

	// CT Logs
	if r.CTLogs != nil && len(r.CTLogs.SubdomainsDiscovered) > 0 {
		sb.WriteString(fmt.Sprintf("## 4. 🛰️ Discovered Subdomains via CT Logs (%d found)\n\n", r.CTLogs.TotalFound))
		for _, sub := range r.CTLogs.SubdomainsDiscovered {
			sb.WriteString(fmt.Sprintf("- `%s`\n", sub))
		}
		sb.WriteString("\n")
	}

	sb.WriteString("---\n*Generated by GopherRecon - Passive OSINT Engine*\n")
	return os.WriteFile(filePath, []byte(sb.String()), 0644)
}
