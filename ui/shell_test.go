package ui_test

import (
	"strings"
	"testing"

	"github.com/mbriggs/gesso/ui"
)

func TestShellCarriesDrawerHooks(t *testing.T) {
	got := renderComponent(t, ui.Shell(ui.ShellProps{
		Sidebar: ui.Components(
			ui.ShellSidebarHeader(ui.ShellSidebarHeaderProps{
				Brand:   ui.ShellBrand(ui.ShellBrandProps{Prefix: "GS", Name: "Studio"}),
				Actions: ui.ThemeToggle(ui.ThemeToggleProps{}),
			}),
			ui.ShellNav(ui.Components(
				ui.ShellNavSection("Catalog"),
				ui.ShellNavItem(ui.ShellNavItemProps{Href: "/families", Label: "Families", Icon: ui.Icon(ui.IconStack, ""), Active: true}),
				ui.ShellNavItem(ui.ShellNavItemProps{Href: "/products", Label: "Products"}),
			)),
			ui.ShellSidebarFooter(ui.Text("footer")),
		),
		MobileHeader: ui.ShellMobileHeader(ui.ShellMobileHeaderProps{}),
		Content:      ui.Text("content"),
	}))
	for _, want := range []string{
		`<aside class="shell-sidebar" data-ui-shell-sidebar>`,
		`<div class="shell-scrim" data-ui-shell-scrim>`,
		`data-ui-shell-toggle`,
		`aria-expanded="false"`,
		`aria-label="Open navigation"`,
		`<span class="shell-sidebar-brand-prefix">GS</span>`,
		`<div class="shell-sidebar-section-label">Catalog</div>`,
		`data-active="true"`,
		`aria-current="page"`,
		`class="shell-nav-icon"`,
		`data-ui-theme-toggle`,
		`class="theme-toggle-moon"`,
		`class="theme-toggle-sun"`,
		`class="shell-sidebar-footer"`,
		`class="shell-content-inner"`,
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("shell missing %q:\n%s", want, got)
		}
	}
	if strings.Count(got, `aria-current="page"`) != 1 {
		t.Fatalf("only the active nav item should be current:\n%s", got)
	}
}
