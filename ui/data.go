package ui

import (
	"strconv"

	"github.com/a-h/templ"
)

type TableFrameProps struct {
	ScrollY      bool
	StickyHeader bool
	MaxHeight    string
	Class        string
	Content      templ.Component
}

type TableVariant string

const (
	TableVariantDefault TableVariant = "default"
	TableVariantData    TableVariant = "data"
)

type TableDensity string

const (
	TableDensityDefault TableDensity = "default"
	TableDensityCompact TableDensity = "compact"
)

type TableLayout string

const (
	TableLayoutAuto  TableLayout = "auto"
	TableLayoutFixed TableLayout = "fixed"
)

type TableProps struct {
	Variant TableVariant
	Density TableDensity
	Layout  TableLayout
	Wrap    bool
	Class   string
	Content templ.Component
	Attrs   templ.Attributes
}

type PaginationProps struct {
	Page     int
	Pages    int
	HrefBase string // e.g. "/products?page=" — page number is appended as last token
	HrefFn   func(page int) string
	Info     templ.Component
	InfoText string
	Window   int
	NoBorder bool
	Class    string
}

type PaginationEntry struct {
	Number int
	Gap    bool
}

type TooltipPosition string

const (
	TooltipTop    TooltipPosition = "top"
	TooltipBottom TooltipPosition = "bottom"
	TooltipLeft   TooltipPosition = "left"
	TooltipRight  TooltipPosition = "right"
)

type TooltipProps struct {
	Text           templ.Component
	TextString     string
	Position       TooltipPosition
	Inline         bool
	NoInline       bool
	FocusableChild bool
	DescribedByID  string
	Trigger        templ.Component
	TriggerText    string
	Class          string
}

type StatusToggleProps struct {
	Name     string
	Value    string
	Checked  bool
	Disabled bool
	OnLabel  string
	OffLabel string
	Class    string
	Attrs    templ.Attributes
}

type FloatingActionBarTone string

const (
	FloatingActionBarToneDefault FloatingActionBarTone = "default"
	FloatingActionBarToneWarning FloatingActionBarTone = "warning"
)

type FloatingActionBarProps struct {
	Tone        FloatingActionBarTone
	NoSticky    bool
	Class       string
	Message     templ.Component
	MessageText string
	Actions     templ.Component
}

func tableFrameClasses(p TableFrameProps) string {
	return Class(
		"table-frame",
		ClassIf(p.ScrollY, "scroll-y"),
		ClassIf(p.StickyHeader, "sticky-header"),
		p.Class,
	)
}

func tableClasses(p TableProps) string {
	return Class(
		"table",
		ClassIf(p.Variant == TableVariantData, "data"),
		ClassIf(p.Density == TableDensityCompact, "compact"),
		ClassIf(p.Layout == TableLayoutFixed, "fixed"),
		ClassIf(p.Wrap, "wrap"),
		p.Class,
	)
}

func paginationClasses(p PaginationProps) string {
	return Class("pagination", ClassIf(!p.NoBorder, "bordered"), p.Class)
}

func paginationWindow(p PaginationProps) int {
	if p.Window <= 0 {
		return 1
	}
	return p.Window
}

func paginationHref(p PaginationProps, page int) templ.SafeURL {
	if p.HrefFn != nil {
		return templ.SafeURL(p.HrefFn(page))
	}
	return templ.SafeURL(p.HrefBase + strconv.Itoa(page))
}

func paginationLabel(n int) string { return strconv.Itoa(n) }

// PaginationSeries returns the visible page entries: always page 1 and
// p.Pages, the current page plus `window` either side, with gaps inserted
// where the sequence jumps.
func PaginationSeries(page, pages, window int) []PaginationEntry {
	if pages <= 1 {
		return nil
	}
	if window < 0 {
		window = 0
	}
	seen := map[int]bool{}
	numbers := make([]int, 0, 8)
	push := func(n int) {
		if n < 1 || n > pages || seen[n] {
			return
		}
		seen[n] = true
		numbers = append(numbers, n)
	}
	push(1)
	for i := page - window; i <= page+window; i++ {
		push(i)
	}
	push(pages)
	// numbers are unique but unordered relative to the canonical page order —
	// sort by value before computing gaps.
	for i := 1; i < len(numbers); i++ {
		for j := i; j > 0 && numbers[j-1] > numbers[j]; j-- {
			numbers[j-1], numbers[j] = numbers[j], numbers[j-1]
		}
	}
	out := make([]PaginationEntry, 0, len(numbers)+2)
	for i, n := range numbers {
		if i > 0 && n-numbers[i-1] > 1 {
			out = append(out, PaginationEntry{Gap: true})
		}
		out = append(out, PaginationEntry{Number: n})
	}
	return out
}

func tooltipClasses(p TooltipProps) string {
	inline := !p.NoInline
	return Class("tooltip", ClassIf(inline, "inline"), p.Class)
}

func tooltipPosition(p TooltipProps) string {
	if p.Position == "" {
		return string(TooltipTop)
	}
	return string(p.Position)
}

func statusToggleClasses(p StatusToggleProps) string {
	state := "off"
	if p.Checked {
		state = "on"
	}
	return Class("status-toggle", state, ClassIf(p.Disabled, "disabled"), p.Class)
}

func statusToggleLabel(p StatusToggleProps) string {
	if p.Checked {
		return stringDefault(p.OnLabel, "On")
	}
	return stringDefault(p.OffLabel, "Off")
}

func floatingActionBarClasses(p FloatingActionBarProps) string {
	tone := string(p.Tone)
	if tone == "" {
		tone = string(FloatingActionBarToneDefault)
	}
	return Class("floating-action-bar", tone, ClassIf(!p.NoSticky, "sticky"), p.Class)
}
