package banner_test

import (
	"testing"

	"github.com/CYBERWOLF333/gopherrecon/pkg/banner"
)

func TestColorHelpers(t *testing.T) {
	red := banner.ColorRed("alert")
	if red == "alert" {
		t.Errorf("expected ANSI wrapped red string")
	}

	green := banner.ColorGreen("ok")
	if green == "ok" {
		t.Errorf("expected ANSI wrapped green string")
	}
}
