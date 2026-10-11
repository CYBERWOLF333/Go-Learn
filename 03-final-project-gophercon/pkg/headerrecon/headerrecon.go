package headerrecon

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

type HeaderCheck struct {
	HeaderName string `json:"header_name"`
	Present    bool   `json:"present"`
	Value      string `json:"value,omitempty"`
	Weight     int    `json:"weight"`
	Severity   string `json:"severity"`
	Remedy     string `json:"remedy"`
}

type HeaderReport struct {
	URL              string        `json:"url"`
	StatusCode       int           `json:"status_code"`
	SecurityScore    int           `json:"security_score"`
	Grade            string        `json:"grade"`
	Checks           []HeaderCheck `json:"checks"`
	ServerBanner     string        `json:"server_banner,omitempty"`
	PoweredByBanner  string        `json:"powered_by_banner,omitempty"`
	InformationLeaks []string      `json:"information_leaks,omitempty"`
	Error            string        `json:"error,omitempty"`
}

var standardRules = []struct {
	name     string
	weight   int
	severity string
	remedy   string
}{
	{"Strict-Transport-Security", 25, "HIGH", "Add: Strict-Transport-Security: max-age=31536000; includeSubDomains"},
	{"Content-Security-Policy", 25, "HIGH", "Add: Content-Security-Policy: default-src 'self'"},
	{"X-Frame-Options", 15, "MEDIUM", "Add: X-Frame-Options: SAMEORIGIN"},
	{"X-Content-Type-Options", 15, "MEDIUM", "Add: X-Content-Type-Options: nosniff"},
	{"Referrer-Policy", 10, "LOW", "Add: Referrer-Policy: strict-origin-when-cross-origin"},
	{"Permissions-Policy", 10, "LOW", "Add: Permissions-Policy: camera=(), microphone=(), geolocation=()"},
}

func calculateGrade(score int) string {
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

// AnalyzeHeaders connects to the target web endpoint and audits headers
func AnalyzeHeaders(domain string, timeout time.Duration) *HeaderReport {
	report := &HeaderReport{
		URL: "https://" + domain,
	}

	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 3 {
				return http.ErrUseLastResponse
			}
			return nil
		},
	}

	req, err := http.NewRequest("GET", report.URL, nil)
	if err != nil {
		report.Error = fmt.Sprintf("Request build error: %v", err)
		return report
	}
	req.Header.Set("User-Agent", "GopherRecon-OSINT-Engine/1.0")

	resp, err := client.Do(req)
	if err != nil {
		// Fallback to http if https failed
		report.URL = "http://" + domain
		req.URL.Scheme = "http"
		resp, err = client.Do(req)
		if err != nil {
			report.Error = fmt.Sprintf("HTTP connection failed: %v", err)
			return report
		}
	}
	defer resp.Body.Close()

	report.StatusCode = resp.StatusCode
	report.ServerBanner = resp.Header.Get("Server")
	report.PoweredByBanner = resp.Header.Get("X-Powered-By")

	if report.ServerBanner != "" {
		report.InformationLeaks = append(report.InformationLeaks, "Server: "+report.ServerBanner)
	}
	if report.PoweredByBanner != "" {
		report.InformationLeaks = append(report.InformationLeaks, "X-Powered-By: "+report.PoweredByBanner)
	}

	totalScore := 0
	for _, r := range standardRules {
		val := resp.Header.Get(r.name)
		if val != "" {
			totalScore += r.weight
			report.Checks = append(report.Checks, HeaderCheck{
				HeaderName: r.name,
				Present:    true,
				Value:      val,
				Weight:     r.weight,
				Severity:   r.severity,
			})
		} else {
			report.Checks = append(report.Checks, HeaderCheck{
				HeaderName: r.name,
				Present:    false,
				Weight:     r.weight,
				Severity:   r.severity,
				Remedy:     r.remedy,
			})
		}
	}

	report.SecurityScore = totalScore
	report.Grade = calculateGrade(totalScore)
	return report
}
