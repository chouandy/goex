package stringsex

import (
	"strings"
	"testing"
)

func TestParseMultipleLineToStringArray(t *testing.T) {
	in := strings.Join([]string{
		"line1",
		"line2",
		"",
		"line3",
		"",
	}, "\n")

	want := []string{
		"line1",
		"line2",
		"line3",
	}

	got := ParseMultipleLineToStringArray(in)

	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Errorf("ParseMultipleLineToStringArray(%q) == %q, want %q", in, got, want)
	}
}
