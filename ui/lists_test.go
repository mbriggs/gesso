package ui_test

import (
	"strings"
	"testing"

	"github.com/mbriggs/gesso/ui"
)

func TestListRowsLinkAndStaticVariants(t *testing.T) {
	got := renderComponent(t, ui.ListGroup(ui.ListGroupProps{Content: ui.Components(
		ui.ListRow(ui.ListRowProps{
			Href:     "/families/42",
			Leading:  ui.Icon(ui.IconFolder, ""),
			Title:    "Business cards",
			MetaText: "4 devices",
			Trailing: ui.Badge(ui.BadgeProps{Tone: ui.ToneSuccess, Text: "ready"}),
		}),
		ui.ListRow(ui.ListRowProps{Title: "Stickers"}),
	)}))
	for _, want := range []string{
		`<ul class="list-group">`,
		`<a class="list-group-link" href="/families/42">`,
		`class="list-group-title"`,
		`<div class="list-group-meta">4 devices</div>`,
		`class="list-group-trailing"`,
		`class="list-group-chevron"`,
		`<div class="list-group-link">`,
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("list group missing %q:\n%s", want, got)
		}
	}
	if strings.Count(got, "list-group-chevron") != 1 {
		t.Fatalf("only the linked row should carry a chevron:\n%s", got)
	}
}
