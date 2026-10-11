package main

import "fmt"

// 1. Define an Interface (The USB Plug Standard)
type SecurityAuditor interface {
	Audit(target string) string
	AuditorName() string
}

// 2. Implementation A: DNS Auditor
type DNSAuditor struct {
	Name string
}

func (d DNSAuditor) AuditorName() string {
	return d.Name
}

func (d DNSAuditor) Audit(target string) string {
	return fmt.Sprintf("DNS Audit for %s: Found A and MX records. No zone transfer allowed. ✅", target)
}

// 3. Implementation B: TLS Certificate Auditor
type TLSAuditor struct {
	Name         string
	MinKeyLength int
}

func (t TLSAuditor) AuditorName() string {
	return t.Name
}

func (t TLSAuditor) Audit(target string) string {
	return fmt.Sprintf("TLS Audit for %s: TLS 1.3 active with RSA-%d key. Cert valid for 82 days. 🔒",
		target, t.MinKeyLength)
}

// 4. A Universal Runner function that accepts ANY SecurityAuditor!
func RunSecurityCheck(auditor SecurityAuditor, target string) {
	fmt.Printf("[ENGINE] Launching auditor '%s'...\n", auditor.AuditorName())
	result := auditor.Audit(target)
	fmt.Println("  ->", result)
}

func main() {
	fmt.Println("=== 🎖️ Lesson 07: Methods & Interfaces ===")

	dnsTool := DNSAuditor{Name: "Passive-DNS-Gopher"}
	tlsTool := TLSAuditor{Name: "Crypto-Cert-Hawk", MinKeyLength: 4096}

	target := "example.com"

	// We can put both tools into a slice of SecurityAuditor because both satisfy the interface!
	pipeline := []SecurityAuditor{dnsTool, tlsTool}

	fmt.Printf("Running Security Pipeline on target: %s\n\n", target)
	for _, auditor := range pipeline {
		RunSecurityCheck(auditor, target)
	}
}
