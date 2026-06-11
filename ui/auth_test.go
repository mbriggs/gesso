package ui_test

import (
	"strings"
	"testing"

	"github.com/mbriggs/gesso/ui"
)

func TestAuthSurfaceComposesCardSlots(t *testing.T) {
	got := renderComponent(t, ui.AuthPage(ui.AuthPageProps{
		Content: ui.AuthCard(ui.AuthCardProps{
			Brand:       "GS",
			Headline:    "Sign in",
			Subheadline: "Use your admin account to continue.",
			Flash:       ui.AuthFlash(ui.AuthFlashProps{Text: "Try another email address or password."}),
			Content:     ui.AuthFields(ui.Text("fields")),
			Footer:      ui.Text("Forgot password?"),
		}),
	}))
	for _, want := range []string{
		`class="auth-page"`,
		`class="auth-glow indigo"`,
		`class="auth-glow cyan"`,
		`class="auth-card"`,
		`<span class="brand-mark glow" aria-hidden="true">GS</span>`,
		`<h1 class="auth-headline">Sign in</h1>`,
		`<p class="auth-subheadline">Use your admin account to continue.</p>`,
		`class="auth-card-flash" role="alert"`,
		`<div class="auth-field-stack">`,
		`<div class="auth-card-footer">`,
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("auth surface missing %q:\n%s", want, got)
		}
	}

	ok := renderComponent(t, ui.AuthFlash(ui.AuthFlashProps{OK: true, Text: "Signed out."}))
	if !strings.Contains(ok, `class="auth-card-flash ok" role="status"`) {
		t.Fatalf("ok flash should be status-toned:\n%s", ok)
	}
}
