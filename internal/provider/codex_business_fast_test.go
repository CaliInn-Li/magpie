package provider

import (
	"testing"
	"time"
)

func TestBusinessFastActive(t *testing.T) {
	before := BusinessFastUntil.Add(-time.Second)
	after := BusinessFastUntil.Add(time.Second)
	cases := []struct {
		name   string
		now    time.Time
		logins []Login
		want   bool
	}{
		{"active business", before, []Login{{Agent: "codex", Plan: "Business", Active: true}}, true},
		{"on business", before, []Login{{Agent: "codex", Plan: "business", On: true}}, true},
		{"inactive business", before, []Login{{Agent: "codex", Plan: "Business"}}, false},
		{"non business", before, []Login{{Agent: "codex", Plan: "Plus", Active: true}}, false},
		{"unknown plan", before, []Login{{Agent: "codex", Active: true}}, false},
		{"after promotion", after, []Login{{Agent: "codex", Plan: "Business", Active: true}}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := businessFastActive(tc.now, tc.logins); got != tc.want {
				t.Fatalf("businessFastActive() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestBusinessFastDeadlineOverride(t *testing.T) {
	t.Setenv("MAGPIE_CODEX_BUSINESS_FAST_UNTIL", "2027-01-15T00:00:00Z")
	if got := businessFastDeadline(); !got.Equal(time.Date(2027, 1, 15, 0, 0, 0, 0, time.UTC)) {
		t.Fatalf("businessFastDeadline() = %v", got)
	}
	t.Setenv("MAGPIE_CODEX_BUSINESS_FAST_UNTIL", "not-a-date")
	if got := businessFastDeadline(); !got.Equal(BusinessFastUntil) {
		t.Fatalf("invalid override = %v, want %v", got, BusinessFastUntil)
	}
}
