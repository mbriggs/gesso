package ui

type ComboboxChoice struct {
	Label string
	Value string
}

type ComboboxProps struct {
	ID           string
	Name         string
	Value        string
	Choices      []ComboboxChoice
	Label        string
	AriaLabel    string
	Hint         string
	Error        string
	Half         bool
	Placeholder  string
	IncludeBlank string
	EmptyMessage string
	Required     bool
	Disabled     bool
	DescribedBy  string
	Class        string
}

func comboboxDescribedBy(p ComboboxProps) string {
	return describedBy(p.ID, p.DescribedBy, p.Hint, p.Error)
}

func comboboxSelectedLabel(p ComboboxProps) string {
	for _, c := range p.Choices {
		if c.Value == p.Value {
			return c.Label
		}
	}
	return ""
}
