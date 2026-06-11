package ui

import "github.com/a-h/templ"

type ListGroupProps struct {
	Class   string
	Content templ.Component // ListRow children
}

type ListRowProps struct {
	// Href makes the whole row a link and appends a trailing chevron.
	// Leave empty for a static row.
	Href         string
	Class        string
	Leading      templ.Component // icon, image, or status mark before the text
	Title        string
	TitleContent templ.Component
	MetaText     string
	Meta         templ.Component
	Trailing     templ.Component // badges or actions before the chevron
}
