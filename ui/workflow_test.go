package ui_test

import (
	"testing"
	"time"

	"github.com/mbriggs/gesso/ui"
)

func TestFormatElapsedBetween(t *testing.T) {
	t.Parallel()
	start := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	cases := []struct {
		name string
		end  time.Time
		want string
	}{
		{"zero is 0s", start, "0s"},
		{"sub-minute", start.Add(45 * time.Second), "45s"},
		{"minute and seconds", start.Add(2*time.Minute + 7*time.Second), "2m 7s"},
		{"hour minute seconds", start.Add(3*time.Hour + 4*time.Minute + 5*time.Second), "3h 4m 5s"},
		{"negative durations clamp to zero", start.Add(-time.Second), "0s"},
		{"exact hour shows 0m 0s suffix", start.Add(time.Hour), "1h 0m 0s"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ui.FormatElapsedBetween(start, tc.end)
			if got != tc.want {
				t.Fatalf("FormatElapsedBetween(%v, %v) = %q, want %q", start, tc.end, got, tc.want)
			}
		})
	}
}

func TestStepperUsesProvidedNow(t *testing.T) {
	t.Parallel()
	start := time.Date(2026, 5, 1, 9, 0, 0, 0, time.UTC)
	now := start.Add(75 * time.Second)
	got := renderComponent(t, ui.Stepper(ui.StepperProps{
		Header: ui.StepperHeader{Title: "Build"},
		Now:    now,
		Steps: []ui.StepperStep{
			{Name: "compile", Label: "Compile", Status: ui.StatusActive, StartedAt: start},
		},
	}))
	if !contains(got, "1m 15s") {
		t.Fatalf("expected elapsed '1m 15s' in stepper output:\n%s", got)
	}
}

func contains(haystack, needle string) bool {
	return len(needle) > 0 && len(haystack) >= len(needle) && indexOf(haystack, needle) >= 0
}

func indexOf(s, sub string) int {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}
