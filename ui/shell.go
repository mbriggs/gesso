package ui

import (
	"strings"

	"github.com/a-h/templ"
)

// ShellProps is the app frame: a sidebar (off-canvas drawer below the
// container breakpoint, static column above it), a scrim, and the scrolling
// content column. The shell is the page's scroll container — the document
// body stays fixed (base.css) so the sidebar never scrolls away.
type ShellProps struct {
	Class        string
	Sidebar      templ.Component // ShellSidebarHeader / ShellNav / ShellSidebarFooter
	MobileHeader templ.Component // ShellMobileHeader; hidden at desktop width
	Content      templ.Component
}

type ShellSidebarHeaderProps struct {
	Brand   templ.Component // typically ShellBrand
	Actions templ.Component // typically ThemeToggle
}

type ShellBrandProps struct {
	Href   string // default "/"
	Prefix string // accent-colored lead-in, e.g. "GS"
	Name   string
}

type ShellNavItemProps struct {
	Href   string
	Label  string
	Icon   templ.Component
	Active bool // data-active="true" + aria-current="page"
	Attrs  templ.Attributes
}

type ShellMobileHeaderProps struct {
	Brand     templ.Component
	Actions   templ.Component // start cluster, next to the brand (typically ThemeToggle)
	End       templ.Component // end cluster (typically ShellSidebarProject + ShellSidebarUserMenu)
	MenuLabel string          // aria-label for the drawer toggle, default "Open navigation"
}

// ShellSidebarProjectProps is the footer's context chip: a status dot plus
// the active account/workspace name. Set Href to render it as a link.
type ShellSidebarProjectProps struct {
	Name string
	Href string
}

// ShellSidebarUserMenuProps is the footer's account affordance: an avatar
// button opening a flyout with the signed-in email and the caller's items.
type ShellSidebarUserMenuProps struct {
	Email   string
	Initial string          // avatar letter; defaults to Email's first letter
	Content templ.Component // menu items rendered below the email banner
}

// avatarInitial picks the single letter shown in the avatar circle.
func avatarInitial(p ShellSidebarUserMenuProps) string {
	if p.Initial != "" {
		return p.Initial
	}
	for _, r := range p.Email {
		return strings.ToUpper(string(r))
	}
	return "?"
}

type ThemeToggleProps struct {
	Class string
}
