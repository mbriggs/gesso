package ui

import "github.com/a-h/templ"

type AlertProps struct {
	Tone    Tone
	Role    string
	Class   string
	Icon    templ.Component
	Title   string
	Content templ.Component
	Text    string
}

type BadgeProps struct {
	Tone    Tone
	Class   string
	Pill    bool
	Content templ.Component
	Text    string
}

type SpinnerProps struct {
	Label string
	Class string
}

type ProgressProps struct {
	Value  int
	Total  int
	Tone   Tone
	Size   string
	Class  string
	Label  string
	Inline bool
}
