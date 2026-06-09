package ui

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

type StepStatus string

const (
	StatusCompleted StepStatus = "completed"
	StatusActive    StepStatus = "active"
	StatusPending   StepStatus = "pending"
	StatusFailed    StepStatus = "failed"
)

type StageBarVariant string

const (
	StageBarDefault  StageBarVariant = "default"
	StageBarEmbedded StageBarVariant = "embedded"
)

type Stage struct {
	Key    string
	Label  string
	Hint   string
	Status StepStatus
	Href   string
}

type StageBarProps struct {
	Stages  []Stage
	Current string
	Label   string
	Variant StageBarVariant
	Class   string
}

type StatusCircleSize string

const (
	StatusCircleSm StatusCircleSize = "sm"
	StatusCircleMd StatusCircleSize = "md"
)

type StatusCircleProps struct {
	Status StepStatus
	Size   StatusCircleSize
	Class  string
}

type StepperHeaderStatus string

const (
	StepperStatusCompleted StepperHeaderStatus = "completed"
	StepperStatusActive    StepperHeaderStatus = "active"
	StepperStatusPending   StepperHeaderStatus = "pending"
	StepperStatusFailed    StepperHeaderStatus = "failed"
	StepperStatusPaused    StepperHeaderStatus = "paused"
)

type StepperHeader struct {
	Title        string
	Status       StepperHeaderStatus
	ElapsedSince time.Time
	IconColor    StepStatus
	Animated     bool
	Icon         templ.Component
	Action       templ.Component
}

type StepperStep struct {
	Name        string
	Label       string
	Status      StepStatus
	Summary     string
	StartedAt   time.Time
	FinishedAt  time.Time
	DetailHref  string
	DetailLabel string
	Children    []StepperStep
}

type StepperProps struct {
	Header StepperHeader
	Steps  []StepperStep
	// Now is the reference clock used to format elapsed time for active
	// steps. Set it explicitly at the handler/page level — leaving it
	// zero falls back to time.Now() at render, which makes tests and
	// hydration non-deterministic.
	Now          time.Time
	Error        string
	ErrorContent templ.Component
	Class        string
}

type EditableGridProps struct {
	DirtyCount   int
	Submitting   bool
	Toolbar      templ.Component
	Error        string
	ErrorContent templ.Component
	DirtyMessage templ.Component
	ResetLabel   string
	ResetAttrs   templ.Attributes
	SaveLabel    string
	Content      templ.Component
	Class        string
}

func stepperNow(p StepperProps) time.Time {
	if !p.Now.IsZero() {
		return p.Now
	}
	return time.Now()
}

func stageBarClasses(p StageBarProps) string {
	variant := string(p.Variant)
	if variant == "" {
		variant = string(StageBarDefault)
	}
	return Class("stage-bar", variant, p.Class)
}

func stageStatusClass(s StepStatus) string {
	if s == StatusPending {
		return ""
	}
	return string(s)
}

func statusCircleClasses(p StatusCircleProps) string {
	size := p.Size
	if size == "" {
		size = StatusCircleMd
	}
	return Class("status-circle", string(size), string(p.Status), p.Class)
}

func statusCircleSize(p StatusCircleProps) string {
	if p.Size == "" {
		return string(StatusCircleMd)
	}
	return string(p.Size)
}

func stepperClasses(p StepperProps) string {
	return Class("card", "shadow", "stepper", p.Class)
}

func stepperLineState(s StepStatus) string {
	if s == StatusCompleted || s == StatusFailed {
		return string(s)
	}
	return ""
}

func editableGridSaveLabel(p EditableGridProps) string {
	if p.SaveLabel != "" {
		return p.SaveLabel
	}
	if p.Submitting {
		return "Saving..."
	}
	if p.DirtyCount == 1 {
		return "Save 1 change"
	}
	return fmt.Sprintf("Save %d changes", p.DirtyCount)
}

func ternaryString(cond bool, a, b string) string {
	if cond {
		return a
	}
	return b
}

func intToString(n int) string {
	return strconv.Itoa(n)
}

func pluralUnsavedChange(n int) string {
	if n == 1 {
		return "unsaved change"
	}
	return "unsaved changes"
}

// FormatElapsedBetween returns a humanized "Xh Ym Zs" duration. Callers pass
// the clock explicitly (time.Now() or a fixed time) so renders stay stable in
// tests.
func FormatElapsedBetween(from, to time.Time) string {
	secs := int64(math.Max(0, math.Floor(to.Sub(from).Seconds())))
	hours := secs / 3600
	minutes := (secs % 3600) / 60
	s := secs % 60
	if hours > 0 {
		return fmt.Sprintf("%dh %dm %ds", hours, minutes, s)
	}
	if minutes > 0 {
		return fmt.Sprintf("%dm %ds", minutes, s)
	}
	return fmt.Sprintf("%ds", s)
}
