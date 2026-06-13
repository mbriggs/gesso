package ui_test

import (
	"bytes"
	"context"
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

func TestPageHeaderRendersFlashFromContext(t *testing.T) {
	ctx := ui.WithFlash(context.Background(), []ui.FlashMessage{
		{Kind: ui.FlashError, Message: "2 active device(s) reference this value"},
	})
	var b bytes.Buffer
	if err := ui.PageHeader(ui.PageHeaderProps{Title: "Acer"}).Render(ctx, &b); err != nil {
		t.Fatalf("render page header: %v", err)
	}
	got := b.String()

	flashAt := strings.Index(got, `<div class="flash">`)
	if flashAt == -1 {
		t.Fatalf("page header should render the ctx flash:\n%s", got)
	}
	if sepAt := strings.Index(got, "page-header-separator"); flashAt < sepAt {
		t.Fatalf("flash should render below the header, not inside it:\n%s", got)
	}

	plain := renderComponent(t, ui.PageHeader(ui.PageHeaderProps{Title: "Acer"}))
	if strings.Contains(plain, `class="flash"`) {
		t.Fatalf("page header without ctx flash should render no flash region:\n%s", plain)
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
