package ui

import "github.com/a-h/templ"

type PageHeaderProps struct {
	Title       string
	Sub         string
	Class       string
	Breadcrumbs templ.Component
	Eyebrow     templ.Component
	Badges      templ.Component
	Meta        templ.Component // datum row under the title (PageDatum et al.)
	Actions     templ.Component
	NoSeparator bool // suppress the full-bleed rule under the header
}

type BreadcrumbItem struct {
	Label   string
	Href    string // empty for the current/last crumb
	Current bool   // true marks the last crumb (aria-current="page")
}

type BreadcrumbsProps struct {
	Items []BreadcrumbItem
	Class string
}

type EmptyStateProps struct {
	Title   string
	Message string
	Spacing string
	Class   string
	Icon    templ.Component
	Action  templ.Component
}

type SectionHeadProps struct {
	Title   string
	Meta    string
	Actions templ.Component
}

type SectionLinkProps struct {
	Href  string
	Text  string
	Class string
}

type EntityTileProps struct {
	ID        string
	Href      string
	Class     string
	WrapClass string
	Muted     bool
	Title     string
	Leading   templ.Component
	Label     templ.Component
	Text      string
	Tag       string
	Meta      templ.Component
	Actions   templ.Component
}
