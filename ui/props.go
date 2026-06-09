package ui

import "strings"

func Class(classes ...string) string {
	out := make([]string, 0, len(classes))
	for _, class := range classes {
		if class != "" {
			out = append(out, class)
		}
	}
	return strings.Join(out, " ")
}

func ClassIf(on bool, className string) string {
	if !on {
		return ""
	}
	return className
}

func describedBy(id, explicit, hint, err string) string {
	if explicit != "" {
		return explicit
	}
	if id == "" {
		return ""
	}
	ids := make([]string, 0, 2)
	if hint != "" {
		ids = append(ids, id+"-hint")
	}
	if err != "" {
		ids = append(ids, id+"-error")
	}
	return strings.Join(ids, " ")
}
