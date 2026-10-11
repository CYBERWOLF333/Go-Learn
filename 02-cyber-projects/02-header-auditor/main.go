package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type SecurityHeaderRule struct {
	HeaderName  string
	Weight      int
	Description string
	Remedy      string
}

type HeaderAuditResult struct {
	TargetURL       string
	TotalScore      int
	Grade           string
	PresentHeaders  map[string]string
	MissingHeaders  []SecurityHeaderRule
	InformationLeak []string
}

// Standard OWASP-recommended security headers rules
var Rules = []SecurityHeaderRule{
	{
		HeaderName:  "Strict-Transport-Security",
		Weight:      25,
		Description: "Enforces secure HTTPS connections and prevents SSL-stripping MITM attacks.",
		Remedy:      "Add: Strict-Transport-Security: max-age=31536000; includeSubDomains",
	},
	{
		HeaderName:  "Content-Security-Policy",
		Weight:      25,
		Description: "Restricts sources from which scripts, images, and objects can be loaded (stops XSS).",
		Remedy:      "Add: Content-Security-Policy: default-src 'self'",
	},
	{
		HeaderName:  "X-Frame-Options",
		Weight:      15,
		Description: "Prevents UI redressing / Clickjacking by disallowing embedding in iframes.",
		Remedy:      "Add: X-Frame-Options: DENY or SAMEORIGIN",
	},
	{
		HeaderName:  "X-Content-Type-Options",
		Weight:      15,
		Description: "Prevents browsers from MIME-sniffing away from the declared content-type.",
		Remedy:      "Add: X-Content-Type-Options: nosniff",
	},
	{
		HeaderName:  "Referrer-Policy",
		Weight:      10,
		Description: "Protects privacy by limiting referrer information leaked in outgoing links.",
		Remedy:      "Add: Referrer-Policy: strict-origin-when-cross-origin",
	},
	{
		HeaderName:  "Permissions-Policy",
		Weight:      10,
		Description: "Restricts browser features like camera, microphone, and geolocation.",
		Remedy:      "Add: Permissions-Policy: geolocation=(), camera=(), microphone=()",
	},
}

func CalculateGrade(score int) string {
	switch {
	case score >= 95:
		return "A+"
	case score >= 85:
		return "A"
	case score >= 75:
		return "B"
	case score >= 60:
		return "C"
	case score >= 40:
		return "D"
	default:
		return "F"
	}
}

func AuditHeaders(headers http.Header, targetURL string) HeaderAuditResult {
	score := 0
	present := make(map[string]string)
	var missing []SecurityHeaderRule
	var infoLeaks []string

	for _, rule := range Rules {
		val := headers.Get(rule.HeaderName)
		if val != "" {
			score += rule.Weight
			present[rule.HeaderName] = val
		} else {
			missing = append(missing, rule)
		}
	}

	// Check for info disclosure headers
	leakHeaders := []string{"Server", "X-Powered-By", "X-AspNet-Version", "X-Generator"}
	for _, leak := range leakHeaders {
		if val := headers.Get(leak); val != "" {
			infoLeaks = append(infoLeaks, fmt.Sprintf("%s: %s", leak, val))
		}
	}

	return HeaderAuditResult{
		TargetURL:       targetURL,
		TotalScore:      score,
		Grade:           CalculateGrade(score),
		PresentHeaders:  present,
		MissingHeaders:  missing,
		InformationLeak: infoLeaks,
	}
}

func FetchAndAudit(targetURL string) (*HeaderAuditResult, error) {
	if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
		targetURL = "https://" + targetURL
	}

	client := &http.Client{
		Timeout: 7 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		},
	}

	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "HeaderAuditor-GopherSecurityBot/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result := AuditHeaders(resp.Header, targetURL)
	return &result, nil
}

func PrintAuditReport(res HeaderAuditResult) {
	fmt.Println("==================================================")
	fmt.Printf(" 🛡️  AUDIT REPORT: %s\n", res.TargetURL)
	fmt.Printf(" 📊 Security Score: %d / 100  |  Grade: [%s]\n", res.TotalScore, res.Grade)
	fmt.Println("==================================================")

	fmt.Println("\n✅ Configured Defensive Headers:")
	if len(res.PresentHeaders) == 0 {
		fmt.Println("   (None! The server is completely unarmored!)")
	} else {
		for k, v := range res.PresentHeaders {
			fmt.Printf("   ✓ %-28s : %s\n", k, v)
		}
	}

	fmt.Println("\n❌ Missing Defensive Headers (Vulnerabilities):")
	if len(res.MissingHeaders) == 0 {
		fmt.Println("   (Awesome! All recommended defensive headers are present!)")
	} else {
		for _, m := range res.MissingHeaders {
			fmt.Printf("   ✗ %-28s (-%d pts)\n", m.HeaderName, m.Weight)
			fmt.Printf("     Why:    %s\n", m.Description)
			fmt.Printf("     Remedy: %s\n", m.Remedy)
		}
	}

	if len(res.InformationLeak) > 0 {
		fmt.Println("\n⚠️  Information Disclosure Warnings (Server Banners):")
		for _, leak := range res.InformationLeak {
			fmt.Printf("   ! Leaking banner -> %s\n", leak)
		}
	}
	fmt.Println()
}

func main() {
	urlFlag := flag.String("url", "", "Target URL to audit (e.g. https://github.com)")
	flag.Parse()

	if *urlFlag == "" {
		fmt.Println("ℹ️  No URL specified. Running Built-in Demo comparison...")

		// Demo 1: Unprotected Server
		badHeaders := http.Header{}
		badHeaders.Set("Server", "Apache/2.4.41 (Ubuntu) mod_ssl/2.4.41")
		badHeaders.Set("X-Powered-By", "PHP/7.4.3")
		resBad := AuditHeaders(badHeaders, "http://vulnerable-mock-app.local")
		PrintAuditReport(resBad)

		// Demo 2: Fortified Server
		goodHeaders := http.Header{}
		goodHeaders.Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		goodHeaders.Set("Content-Security-Policy", "default-src 'self'; script-src 'self'")
		goodHeaders.Set("X-Frame-Options", "DENY")
		goodHeaders.Set("X-Content-Type-Options", "nosniff")
		goodHeaders.Set("Referrer-Policy", "strict-origin-when-cross-origin")
		goodHeaders.Set("Permissions-Policy", "geolocation=()")
		resGood := AuditHeaders(goodHeaders, "https://fortified-vault.local")
		PrintAuditReport(resGood)

		fmt.Println("Tip: Run with '-url https://target.com' to audit any live server!")
		return
	}

	fmt.Printf("[*] Connecting to %s...\n", *urlFlag)
	report, err := FetchAndAudit(*urlFlag)
	if err != nil {
		fmt.Printf("[-] Failed to audit target: %v\n", err)
		return
	}

	PrintAuditReport(*report)
}
