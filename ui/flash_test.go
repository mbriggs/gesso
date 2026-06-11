package ui_test

import (
	"strings"
	"testing"

	"github.com/mbriggs/gesso/ui"
)

func TestFlashMapsKindsToTones(t *testing.T) {
	got := renderComponent(t, ui.Flash(ui.FlashProps{Messages: []ui.FlashMessage{
		{Kind: ui.FlashNotice, Message: "Signed in."},
		{Kind: ui.FlashSuccess, Title: "Saved", Message: "Saved 3 families."},
		{Kind: ui.FlashAlert, Message: "Try another email address or password."},
		{Kind: ui.FlashError, Message: "Something broke."},
	}}))
	for _, want := range []string{
		`<div class="flash">`,
		`class="alert info"`,
		`class="alert success"`,
		`<p class="alert-title">Saved</p>`,
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("flash missing %q:\n%s", want, got)
		}
	}
	if strings.Count(got, `class="alert danger"`) != 2 {
		t.Fatalf("alert and error should both render danger:\n%s", got)
	}

	if empty := renderComponent(t, ui.Flash(ui.FlashProps{})); strings.Contains(empty, "flash") {
		t.Fatalf("empty flash should render nothing:\n%s", empty)
	}
}

func TestFlashToneBridge(t *testing.T) {
	cases := map[ui.FlashKind]ui.Tone{
		ui.FlashNotice:     ui.ToneInfo,
		ui.FlashSuccess:    ui.ToneSuccess,
		ui.FlashAlert:      ui.ToneDanger,
		ui.FlashError:      ui.ToneDanger,
		ui.FlashKind("??"): ui.ToneInfo,
	}
	for kind, want := range cases {
		if got := ui.FlashTone(kind); got != want {
			t.Fatalf("FlashTone(%q) = %q, want %q", kind, got, want)
		}
	}
}
