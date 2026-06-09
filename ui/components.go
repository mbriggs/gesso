package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func Attrs(values ...any) templ.Attributes {
	if len(values)%2 != 0 {
		panic("ui.Attrs expects key/value pairs")
	}
	out := make(templ.Attributes, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			panic("ui.Attrs key must be a string")
		}
		out[key] = values[i+1]
	}
	return out
}

func Components(components ...templ.Component) templ.Component {
	out := make([]templ.Component, 0, len(components))
	for _, component := range components {
		if component != nil {
			out = append(out, component)
		}
	}
	if len(out) == 0 {
		return templ.ComponentFunc(func(context.Context, io.Writer) error { return nil })
	}
	return templ.Join(out...)
}

func stringDefault(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}
