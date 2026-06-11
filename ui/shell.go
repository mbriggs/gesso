package ui

import "github.com/a-h/templ"

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
	Actions   templ.Component
	MenuLabel string // aria-label for the drawer toggle, default "Open navigation"
}

type ThemeToggleProps struct {
	Class string
}
