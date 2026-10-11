package headerrecon_test

import (
	"testing"
	"time"

	"github.com/CYBERWOLF333/gopherrecon/pkg/headerrecon"
)

func TestAnalyzeHeadersOffline(t *testing.T) {
	// Querying non-existent host should fail gracefully without panic
	report := headerrecon.AnalyzeHeaders("127.0.0.1:59999", 100*time.Millisecond)
	if report == nil {
		t.Fatalf("expected non-nil report")
	}
	if report.Error == "" {
		t.Logf("connection completed or errored as expected")
	}
}
