package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateGrade(t *testing.T) {
	tests := []struct {
		score int
		grade string
	}{
		{100, "A+"},
		{95, "A+"},
		{88, "A"},
		{76, "B"},
		{65, "C"},
		{45, "D"},
		{20, "F"},
		{0, "F"},
	}

	for _, tc := range tests {
		got := CalculateGrade(tc.score)
		if got != tc.grade {
			t.Errorf("for score %d expected grade %s, got %s", tc.score, tc.grade, got)
		}
	}
}

func TestAuditHeadersMockServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Server", "SecretDaemon/1.0")
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	res, err := FetchAndAudit(ts.URL)
	if err != nil {
		t.Fatalf("FetchAndAudit failed: %v", err)
	}

	// 25 (HSTS) + 15 (nosniff) = 40
	if res.TotalScore != 40 {
		t.Errorf("expected score 40, got %d", res.TotalScore)
	}

	if len(res.InformationLeak) != 1 {
		t.Errorf("expected 1 information leak header (Server), got %d", len(res.InformationLeak))
	}
}
