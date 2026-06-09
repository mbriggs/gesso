package ui

import "github.com/a-h/templ"

type ButtonProps struct {
	ID           string
	Name         string
	Form         string
	Value        string
	Type         string
	Variant      string
	Size         string
	Class        string
	NoShadow     bool
	Block        bool
	Disabled     bool
	Loading      bool
	LoadingLabel string
	IconLeft     templ.Component
	IconRight    templ.Component
	Content      templ.Component
	Text         string
	Attrs        templ.Attributes
}

type LinkButtonProps struct {
	ID        string
	Href      string
	Variant   string
	Size      string
	Class     string
	NoShadow  bool
	Block     bool
	IconLeft  templ.Component
	IconRight templ.Component
	Content   templ.Component
	Text      string
	Attrs     templ.Attributes
}
