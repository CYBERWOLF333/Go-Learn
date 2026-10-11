package dnsrecon_test

import (
	"testing"
	"time"

	"github.com/CYBERWOLF333/gopherrecon/pkg/dnsrecon"
)

func TestAnalyzeDNSGraceful(t *testing.T) {
	// Querying non-resolvable domain should return structured errors without crashing
	report := dnsrecon.AnalyzeDNS("invalid-local-fake-domain-12345.corp", 500*time.Millisecond)
	if report == nil {
		t.Fatalf("expected non-nil report")
	}
	if report.SPFStatus == "" {
		t.Errorf("expected SPFStatus to be populated")
	}
	if report.DMARCStatus == "" {
		t.Errorf("expected DMARCStatus to be populated")
	}
}
