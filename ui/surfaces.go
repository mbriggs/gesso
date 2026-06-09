package ui

import "github.com/a-h/templ"

type CardProps struct {
	Shadow   bool
	NoShadow bool
	Clipped  bool
	Padded   bool
	Narrow   bool
	Class    string
	Content  templ.Component
	Attrs    templ.Attributes
}

type CardHeaderProps struct {
	Title    string
	Subtitle string
	Class    string
	Content  templ.Component
}

type CardBodyProps struct {
	Class   string
	Content templ.Component
}

type CardFooterProps struct {
	Class   string
	Content templ.Component
}

type StackGap string

const (
	StackGapSm StackGap = "sm"
	StackGapMd StackGap = "md"
	StackGapLg StackGap = "lg"
	StackGapXl StackGap = "xl"
)

type StackProps struct {
	Gap     StackGap
	Class   string
	Content templ.Component
	Attrs   templ.Attributes
}

type SkeletonProps struct {
	Rows    int
	NoTitle bool
	Class   string
}

type SummaryListProps struct {
	Class   string
	Content templ.Component
}

type SummaryItemProps struct {
	Term        templ.Component
	TermText    string
	Details     templ.Component
	DetailsText string
	Class       string
}

func cardClasses(p CardProps) string {
	// Shadow defaults to true unless NoShadow is set. The Shadow field is for
	// explicit opt-in when callers prefer the positive flag; either flag turning
	// it off wins.
	shadow := !p.NoShadow
	return Class(
		"card",
		ClassIf(shadow, "shadow"),
		ClassIf(p.Clipped, "clipped"),
		ClassIf(p.Padded, "padded"),
		ClassIf(p.Narrow, "narrow"),
		p.Class,
	)
}

func stackClasses(p StackProps) string {
	gap := p.Gap
	if gap == "" {
		gap = StackGapMd
	}
	return Class("stack", "gap-"+string(gap), p.Class)
}

func skeletonRowCount(p SkeletonProps) int {
	if p.Rows == 0 {
		return 2
	}
	if p.Rows < 0 {
		return 0
	}
	return p.Rows
}
