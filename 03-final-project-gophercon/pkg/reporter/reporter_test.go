package reporter_test

import (
	"os"
	"testing"
	"time"

	"github.com/CYBERWOLF333/gopherrecon/pkg/dnsrecon"
	"github.com/CYBERWOLF333/gopherrecon/pkg/reporter"
)

func TestExportJSONAndMarkdown(t *testing.T) {
	r := &reporter.FullReconReport{
		TargetDomain: "test.local",
		ScanTime:     time.Now(),
		Duration:     "100ms",
		DNS: &dnsrecon.DNSReport{
			Domain:    "test.local",
			ARecords:  []string{"127.0.0.1"},
			SPFStatus: "STRICT",
		},
	}

	tmpJSON, err := os.CreateTemp("", "test-*.json")
	if err != nil {
		t.Fatal(err)
	}
	tmpJSON.Close()
	defer os.Remove(tmpJSON.Name())

	if err := reporter.ExportJSON(r, tmpJSON.Name()); err != nil {
		t.Fatalf("ExportJSON failed: %v", err)
	}

	tmpMD, err := os.CreateTemp("", "test-*.md")
	if err != nil {
		t.Fatal(err)
	}
	tmpMD.Close()
	defer os.Remove(tmpMD.Name())

	if err := reporter.ExportMarkdown(r, tmpMD.Name()); err != nil {
		t.Fatalf("ExportMarkdown failed: %v", err)
	}

	mdContent, err := os.ReadFile(tmpMD.Name())
	if err != nil {
		t.Fatal(err)
	}
	if len(mdContent) == 0 {
		t.Errorf("markdown report was empty")
	}
}
