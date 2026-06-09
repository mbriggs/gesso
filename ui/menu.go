package ui

import "github.com/a-h/templ"

// MenuProps controls a popover menu rendered with <details>/<summary>.
// Placement defaults to "bottom end". TriggerLabel is required for
// accessibility (it labels both the trigger and the panel).
type MenuProps struct {
	TriggerLabel   string
	TriggerContent templ.Component // overrides the default "..." kebab
	TriggerClass   string          // overrides the default "kebab-trigger" on <summary>
	Placement      string          // e.g. "bottom end", "top start"
	Class          string          // extra classes on the <details> root
	PanelClass     string          // extra classes on the panel
	Content        templ.Component
}

// MenuItemProps renders a <button> styled as a menu item. Inside a <form>,
// set Type="submit"; for actions that need a separate <form>, set Form to
// the form's id.
type MenuItemProps struct {
	Text   string
	Type   string // "button" (default) or "submit"
	Form   string
	Name   string
	Value  string
	Class  string
	Danger bool
	Attrs  templ.Attributes
}

// MenuConfirmProps is the two-click confirmation pattern used inside menus.
// Renders as a <details>: click Label, then Confirm to submit the enclosing
// form.
type MenuConfirmProps struct {
	Label   string // initial menu-item label (default "Continue")
	Confirm string // confirm button label (default Label)
	Message string // optional explanatory text shown after the first click
	Danger  bool
}

func menuPlacement(p string) string {
	if p == "" {
		return "bottom end"
	}
	return p
}

func menuButtonType(t string) string {
	if t == "" {
		return "button"
	}
	return t
}

func menuTriggerClass(c string) string {
	if c == "" {
		return "kebab-trigger"
	}
	return c
}
