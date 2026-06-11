package ui

import (
	"math"
	"slices"
	"strconv"
)

// Tone is the design-system color/severity vocabulary shared by alerts,
// badges, progress bars, and status dots. Use the Tone* constants — empty
// or unrecognized values fall back to a per-component default.
type Tone string

const (
	ToneDefault Tone = "default"
	TonePrimary Tone = "primary"
	ToneInfo    Tone = "info"
	ToneSuccess Tone = "success"
	ToneWarning Tone = "warning"
	ToneDanger  Tone = "danger"
	TonePurple  Tone = "purple"
	ToneHigh    Tone = "high"
)

func buttonClasses(p ButtonProps) string {
	variant := stringDefault(p.Variant, "primary")
	size := stringDefault(p.Size, "md")
	return buttonClass(variant, size, !p.NoShadow, p.Block, p.Class)
}

func linkButtonClasses(p LinkButtonProps) string {
	variant := stringDefault(p.Variant, "primary")
	size := stringDefault(p.Size, "md")
	return buttonClass(variant, size, !p.NoShadow, p.Block, p.Class)
}

// ButtonClasses returns the class string applied by Button/ButtonLink for the
// given variant/size, with shadow defaulted on. Useful when rendering raw
// <button> markup outside the Button templ (e.g. design-system state grids).
func ButtonClasses(variant, size string) string {
	variant = stringDefault(variant, "primary")
	size = stringDefault(size, "md")
	return buttonClass(variant, size, true, false, "")
}

func buttonClass(variant string, size string, shadow bool, block bool, class string) string {
	if variant == "link" {
		return Class("button", "link", ClassIf(block, "block"), class)
	}
	return Class("button", variant, size, ClassIf(shadow && variant != "ghost", "shadow"), ClassIf(block, "block"), class)
}

func inputClasses(p InputProps) string {
	return controlClasses(p.Class, p.Variant, p.Align, p.Invalid || p.Error != "", p.Mono)
}

func inputDescribedBy(p InputProps) string {
	return describedBy(p.ID, p.DescribedBy, p.Hint, p.Error)
}

func selectClasses(p SelectProps) string {
	return controlClasses(p.Class, p.Variant, "", p.Invalid || p.Error != "", false)
}

func selectDescribedBy(p SelectProps) string {
	return describedBy(p.ID, p.DescribedBy, p.Hint, p.Error)
}

func textareaClasses(p TextareaProps) string {
	return controlClasses(p.Class, p.Variant, "", p.Invalid || p.Error != "", false)
}

func textareaDescribedBy(p TextareaProps) string {
	return describedBy(p.ID, p.DescribedBy, p.Hint, p.Error)
}

func checkboxClasses(p CheckboxProps) string {
	return Class("checkbox", p.Class)
}

func toggleFieldClasses(p ToggleProps) string {
	return fieldClasses(p.Class, p.Half)
}

func toggleDescribedBy(p ToggleProps) string {
	return describedBy(p.ID, p.DescribedBy, p.Hint, p.Error)
}

func segmentedFieldClasses(p SegmentedProps) string {
	return fieldClasses(p.Class, p.Half)
}

func segmentedDescribedBy(p SegmentedProps) string {
	return describedBy(p.ID, p.DescribedBy, p.Hint, p.Error)
}

func fieldPropClasses(p FieldProps) string {
	return fieldClasses(p.Class, p.Half)
}

func fieldClasses(class string, half bool) string {
	return Class("field", ClassIf(half, "half"), class)
}

func controlClasses(class, variant, align string, invalid, mono bool) string {
	return Class(
		"input",
		ClassIf(invalid, "invalid"),
		ClassIf(mono, "mono"),
		ClassIf(variant == "table", "table-input"),
		prefixed("align-", align),
		class,
	)
}

func alertClasses(p AlertProps) string {
	tone := coerceTone(p.Tone, []Tone{ToneInfo, ToneSuccess, ToneWarning, ToneDanger}, ToneInfo)
	return Class("alert", string(tone), p.Class)
}

func alertRole(p AlertProps) string {
	if p.Role != "" {
		return p.Role
	}
	if coerceTone(p.Tone, []Tone{ToneDanger}, "") == ToneDanger {
		return "alert"
	}
	return "status"
}

func badgeClasses(p BadgeProps) string {
	tone := coerceTone(p.Tone, []Tone{
		ToneDefault, TonePrimary, ToneInfo, ToneSuccess, ToneWarning, ToneDanger, TonePurple, ToneHigh,
	}, ToneDefault)
	return Class("badge", string(tone), ClassIf(p.Pill, "pill"), p.Class)
}

func progressClasses(p ProgressProps) string {
	return Class("progress", stringDefault(p.Size, "md"), ClassIf(p.Inline, "inline"), p.Class)
}

func progressBarClasses(p ProgressProps) string {
	tone := coerceTone(p.Tone, []Tone{
		ToneDefault, TonePrimary, ToneInfo, ToneSuccess, ToneWarning, ToneDanger,
	}, ToneInfo)
	return Class("progress-bar", string(tone))
}

func progressStyle(value, total int) string {
	return "--progress-value: " + strconv.Itoa(progressPercent(value, total)) + "%"
}

func progressPercent(value, total int) int {
	if total == 0 {
		return 0
	}
	percent := int(math.Round((float64(value) / float64(total)) * 100))
	if percent < 0 {
		return 0
	}
	if percent > 100 {
		return 100
	}
	return percent
}

func statusDotClasses(tone Tone) string {
	tone = coerceTone(tone, []Tone{
		ToneDefault, TonePrimary, ToneInfo, ToneSuccess, ToneWarning, ToneDanger,
	}, ToneDefault)
	return Class("status-dot", string(tone))
}

func prefixed(prefix, value string) string {
	if value == "" {
		return ""
	}
	return prefix + value
}

// coerceTone returns input if it is one of the allowed tones, else fallback.
// Empty input always falls back. No alias mapping — call sites should pass
// the canonical Tone* constants. The flash→tone bridge is FlashTone (flash.go).
func coerceTone(input Tone, allow []Tone, fallback Tone) Tone {
	if input == "" {
		return fallback
	}
	if slices.Contains(allow, input) {
		return input
	}
	return fallback
}
