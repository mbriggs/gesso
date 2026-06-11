package ui_test

import (
	"strings"
	"testing"

	"github.com/mbriggs/gesso/ui"
)

func TestTabsMarkCurrentServerSide(t *testing.T) {
	got := renderComponent(t, ui.Tabs(ui.TabsProps{
		AriaLabel: "Family sections",
		Items: []ui.TabItem{
			{Label: "Overview", Href: "/families/42", Current: true},
			{Label: "Devices", Href: "/families/42/devices", Count: "12"},
		},
	}))
	for _, want := range []string{
		`aria-label="Family sections"`,
		`aria-current="page"`,
		`class="tab-item"`,
		`<span class="tab-count">12</span>`,
		`href="/families/42/devices"`,
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("tabs missing %q:\n%s", want, got)
		}
	}
	if strings.Count(got, `aria-current="page"`) != 1 {
		t.Fatalf("exactly one tab should be current:\n%s", got)
	}

	if empty := renderComponent(t, ui.Tabs(ui.TabsProps{})); strings.Contains(empty, "<nav") {
		t.Fatalf("empty tabs should render nothing:\n%s", empty)
	}
}
