package ui

import "github.com/a-h/templ"

type AuthPageProps struct {
	Class   string
	Content templ.Component // typically one AuthCard
}

type AuthCardProps struct {
	Brand       string // brand-mark glyph text, e.g. "GS"
	Headline    string
	Subheadline string
	Flash       templ.Component // typically AuthFlash; renders between copy and fields
	Content     templ.Component // field stack + submit
	Footer      templ.Component // helper line, e.g. the forgot-password link
	Class       string
}

type AuthFlashProps struct {
	// OK renders the success styling (and role="status"); the default is the
	// failure styling with role="alert".
	OK    bool
	Text  string
	Class string
}

func authFlashRole(p AuthFlashProps) string {
	if p.OK {
		return "status"
	}
	return "alert"
}
