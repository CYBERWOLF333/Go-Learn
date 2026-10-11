package dnsrecon

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"
)

type DNSReport struct {
	Domain      string   `json:"domain"`
	ARecords    []string `json:"a_records"`
	AAAARecords []string `json:"aaaa_records"`
	MXRecords   []string `json:"mx_records"`
	NSRecords   []string `json:"ns_records"`
	TXTRecords  []string `json:"txt_records"`
	SPFRecord   string   `json:"spf_record"`
	SPFStatus   string   `json:"spf_status"`
	DMARCRecord string   `json:"dmarc_record"`
	DMARCStatus string   `json:"dmarc_status"`
	Errors      []string `json:"errors,omitempty"`
}

// AnalyzeDNS gathers DNS intel and evaluates email security posture
func AnalyzeDNS(domain string, timeout time.Duration) *DNSReport {
	report := &DNSReport{
		Domain: domain,
	}

	resolver := &net.Resolver{}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// 1. A Records (IPv4)
	ips, err := resolver.LookupIP(ctx, "ip4", domain)
	if err == nil {
		for _, ip := range ips {
			report.ARecords = append(report.ARecords, ip.String())
		}
	} else {
		report.Errors = append(report.Errors, fmt.Sprintf("A records: %v", err))
	}

	// 2. AAAA Records (IPv6)
	ip6s, err := resolver.LookupIP(ctx, "ip6", domain)
	if err == nil {
		for _, ip := range ip6s {
			report.AAAARecords = append(report.AAAARecords, ip.String())
		}
	}

	// 3. MX Records (Mail Exchangers)
	mxList, err := resolver.LookupMX(ctx, domain)
	if err == nil {
		for _, mx := range mxList {
			report.MXRecords = append(report.MXRecords, fmt.Sprintf("%s (prio %d)", mx.Host, mx.Pref))
		}
	}

	// 4. NS Records (Name Servers)
	nsList, err := resolver.LookupNS(ctx, domain)
	if err == nil {
		for _, ns := range nsList {
			report.NSRecords = append(report.NSRecords, ns.Host)
		}
	}

	// 5. TXT Records & SPF Inspection
	txtList, err := resolver.LookupTXT(ctx, domain)
	if err == nil {
		for _, txt := range txtList {
			report.TXTRecords = append(report.TXTRecords, txt)
			if strings.HasPrefix(txt, "v=spf1") {
				report.SPFRecord = txt
				if strings.Contains(txt, "-all") {
					report.SPFStatus = "STRICT (-all: Hard fail unauthorized senders)"
				} else if strings.Contains(txt, "~all") {
					report.SPFStatus = "MODERATE (~all: Soft fail unauthorized senders)"
				} else if strings.Contains(txt, "+all") || strings.Contains(txt, "?all") {
					report.SPFStatus = "WEAK (?all/+all: Allows spoofing!)"
				} else {
					report.SPFStatus = "CONFIGURED (Custom qualifiers)"
				}
			}
		}
	}

	if report.SPFRecord == "" {
		report.SPFStatus = "MISSING! Domain is vulnerable to email spoofing/phishing!"
	}

	// 6. DMARC Inspection (_dmarc.<domain>)
	dmarcDomain := "_dmarc." + domain
	dmarcTxt, err := resolver.LookupTXT(ctx, dmarcDomain)
	if err == nil && len(dmarcTxt) > 0 {
		for _, txt := range dmarcTxt {
			if strings.HasPrefix(txt, "v=DMARC1") {
				report.DMARCRecord = txt
				if strings.Contains(txt, "p=reject") {
					report.DMARCStatus = "STRONG (p=reject: Spoofed emails dropped)"
				} else if strings.Contains(txt, "p=quarantine") {
					report.DMARCStatus = "MODERATE (p=quarantine: Spoofed emails sent to spam)"
				} else if strings.Contains(txt, "p=none") {
					report.DMARCStatus = "MONITORING ONLY (p=none: Spoofed emails still delivered)"
				} else {
					report.DMARCStatus = "CONFIGURED"
				}
				break
			}
		}
	}

	if report.DMARCRecord == "" {
		report.DMARCStatus = "MISSING! No DMARC protection active."
	}

	return report
}
