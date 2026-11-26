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
