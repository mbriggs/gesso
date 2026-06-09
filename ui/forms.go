package ui

import "github.com/a-h/templ"

type InputProps struct {
	ID           string
	Name         string
	Form         string
	Value        string
	Type         string
	Placeholder  string
	Autocomplete string
	AriaLabel    string
	DescribedBy  string
	Hint         string
	Error        string
	Class        string
	Variant      string
	Align        string
	Required     bool
	Autofocus    bool
	Disabled     bool
	Invalid      bool
	Mono         bool
	Attrs        templ.Attributes
}

type SelectProps struct {
	ID          string
	Name        string
	Form        string
	AriaLabel   string
	DescribedBy string
	Hint        string
	Error       string
	Class       string
	Variant     string
	Required    bool
	Disabled    bool
	Invalid     bool
	Content     templ.Component
	Attrs       templ.Attributes
}

type TextareaProps struct {
	ID          string
	Name        string
	Value       string
	Placeholder string
	AriaLabel   string
	DescribedBy string
	Hint        string
	Error       string
	Class       string
	Variant     string
	Required    bool
	Disabled    bool
	Invalid     bool
	Attrs       templ.Attributes
}

type CheckboxProps struct {
	ID       string
	Name     string
	Form     string
	Value    string
	Label    string
	Class    string
	Checked  bool
	Disabled bool
	Attrs    templ.Attributes
}

type ToggleProps struct {
	ID          string
	Name        string
	Value       string
	Label       string
	Caption     string
	DescribedBy string
	Hint        string
	Error       string
	Class       string
	Half        bool
	Checked     bool
	Disabled    bool
}

type SegmentedChoice struct {
	Label string
	Value string
}

type SegmentedProps struct {
	ID          string
	Name        string
	Value       string
	Label       string
	AriaLabel   string
	DescribedBy string
	Hint        string
	Error       string
	Class       string
	Half        bool
	Choices     []SegmentedChoice
}

type FieldProps struct {
	ID      string
	Label   string
	Control templ.Component
	Hint    string
	Error   string
	Class   string
	Half    bool
}

func segmentedTabIndex(active bool) string {
	if active {
		return "0"
	}
	return "-1"
}

func boolString(value bool) string {
	if value {
		return "true"
	}
	return "false"
}
