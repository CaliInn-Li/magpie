package provider

import (
	"strings"
	"time"
)

// BusinessFastUntil is the end of the temporary ChatGPT Business Fast-mode
// promotion. It is kept as a variable so tests and a future settings surface
// can replace the date without changing routing code.
var BusinessFastUntil = time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)

// BusinessFastActive reports whether a remembered, active Codex account is a
// Business plan account during the promotion window. Unknown plans and
// expired windows are deliberately false.
func BusinessFastActive(now time.Time) bool {
	return businessFastActive(now, Logins("codex"))
}

func businessFastActive(now time.Time, logins []Login) bool {
	if !now.Before(businessFastDeadline()) {
		return false
	}
	for _, login := range logins {
		if !login.Active && !login.On {
			continue
		}
		if strings.Contains(strings.ToLower(strings.TrimSpace(login.Plan)), "business") {
			return true
		}
	}
	return false
}
