package tlsrecon

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

type TLSReport struct {
	Active            bool     `json:"active"`
	Version           string   `json:"version"`
	CipherSuite       string   `json:"cipher_suite"`
	SubjectCN         string   `json:"subject_cn"`
	IssuerOrg         string   `json:"issuer_org"`
	ValidFrom         string   `json:"valid_from"`
	ValidUntil        string   `json:"valid_until"`
	DaysRemaining     int      `json:"days_remaining"`
	IsExpired         bool     `json:"is_expired"`
	AltNames          []string `json:"alt_names"`
	Error             string   `json:"error,omitempty"`
}

func tlsVersionName(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0 (INSECURE / DEPRECATED)"
	case tls.VersionTLS11:
		return "TLS 1.1 (DEPRECATED)"
	case tls.VersionTLS12:
		return "TLS 1.2 (SECURE)"
	case tls.VersionTLS13:
		return "TLS 1.3 (MODERN & RECOMMENDED)"
	default:
		return fmt.Sprintf("Unknown (0x%04x)", version)
	}
}

// AnalyzeTLS establishes a TLS session and extracts cryptographic posture and SANs
func AnalyzeTLS(domain string, timeout time.Duration) *TLSReport {
	report := &TLSReport{}
	target := fmt.Sprintf("%s:443", domain)

	dialer := &net.Dialer{Timeout: timeout}
	conn, err := tls.DialWithDialer(dialer, "tcp", target, &tls.Config{
		ServerName:         domain,
		InsecureSkipVerify: false,
	})
	if err != nil {
		report.Active = false
		report.Error = fmt.Sprintf("TLS Handshake failed: %v", err)
		return report
	}
	defer conn.Close()

	state := conn.ConnectionState()
	report.Active = true
	report.Version = tlsVersionName(state.Version)
	report.CipherSuite = tls.CipherSuiteName(state.CipherSuite)

	if len(state.PeerCertificates) > 0 {
		cert := state.PeerCertificates[0]
		report.SubjectCN = cert.Subject.CommonName
		if len(cert.Issuer.Organization) > 0 {
			report.IssuerOrg = cert.Issuer.Organization[0]
		} else {
			report.IssuerOrg = cert.Issuer.CommonName
		}
		report.ValidFrom = cert.NotBefore.Format("2006-01-02")
		report.ValidUntil = cert.NotAfter.Format("2006-01-02")

		remaining := time.Until(cert.NotAfter)
		report.DaysRemaining = int(remaining.Hours() / 24)
		if report.DaysRemaining < 0 {
			report.IsExpired = true
		}

		report.AltNames = cert.DNSNames
	}

	return report
}
