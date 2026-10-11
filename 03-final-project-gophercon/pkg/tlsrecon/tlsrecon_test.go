package tlsrecon_test

import (
	"testing"
	"time"

	"github.com/CYBERWOLF333/gopherrecon/pkg/tlsrecon"
)

func TestAnalyzeTLSGraceful(t *testing.T) {
	// Querying closed port should return Active=false with error, not panic
	report := tlsrecon.AnalyzeTLS("127.0.0.1:59998", 100*time.Millisecond)
	if report == nil {
		t.Fatalf("expected non-nil report")
	}
	if report.Active {
		t.Errorf("expected inactive TLS for closed port")
	}
	if report.Error == "" {
		t.Errorf("expected error message for closed port")
	}
}
