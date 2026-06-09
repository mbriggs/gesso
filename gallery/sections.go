package gallery

import (
	"time"

	templpkg "github.com/a-h/templ"
)

type Section struct {
	ID    string
	Title string
	Demo  templpkg.Component
}

type SectionGroup struct {
	Title    string
	Sections []Section
}

type PageData struct {
	Groups []SectionGroup
}

func Sections() []SectionGroup {
	return []SectionGroup{
		{
			Title: "Product surfaces",
			Sections: []Section{
				{ID: "product-surface", Title: "Product surface", Demo: designDemoProductSurface()},
				{ID: "page-header", Title: "Page header", Demo: designDemoPageHeader()},
				{ID: "tables", Title: "Tables", Demo: designDemoTables()},
				{ID: "auth", Title: "Auth surface", Demo: designDemoAuth()},
			},
		},
		{
			Title: "Foundations",
			Sections: []Section{
				{ID: "tokens", Title: "Tokens", Demo: designDemoTokens()},
				{ID: "theme-preview", Title: "Theme preview", Demo: designDemoThemePreview()},
				{ID: "typography", Title: "Typography", Demo: designDemoTypography()},
			},
		},
		{
			Title: "Actions",
			Sections: []Section{
				{ID: "buttons", Title: "Buttons", Demo: designDemoButtons()},
				{ID: "button-links", Title: "Link buttons", Demo: designDemoButtonLinks()},
				{ID: "menus", Title: "Menus", Demo: designDemoMenus()},
				{ID: "modal", Title: "Modal", Demo: designDemoModal()},
				{ID: "confirm-dialog", Title: "Confirm dialog", Demo: designDemoConfirmDialog()},
			},
		},
		{
			Title: "Inputs",
			Sections: []Section{
				{ID: "forms", Title: "Forms", Demo: designDemoForms()},
				{ID: "combobox", Title: "Combobox", Demo: designDemoCombobox()},
				{ID: "toggle", Title: "Toggle", Demo: designDemoToggle()},
				{ID: "segmented", Title: "Segmented", Demo: designDemoSegmented()},
			},
		},
		{
			Title: "Navigation",
			Sections: []Section{
				{ID: "pagination", Title: "Pagination", Demo: designDemoPagination()},
			},
		},
		{
			Title: "Feedback",
			Sections: []Section{
				{ID: "alerts", Title: "Alerts", Demo: designDemoAlerts()},
				{ID: "feedback", Title: "Feedback", Demo: designDemoFeedback()},
				{ID: "stage-bar", Title: "Stage bar", Demo: designDemoStageBar()},
				{ID: "stepper", Title: "Stepper", Demo: designDemoStepper()},
			},
		},
		{
			Title: "Supporting",
			Sections: []Section{
				{ID: "badges", Title: "Badges", Demo: designDemoBadges()},
				{ID: "cards", Title: "Cards", Demo: designDemoCards()},
				{ID: "stack", Title: "Stack", Demo: designDemoStack()},
				{ID: "tooltips", Title: "Tooltips", Demo: designDemoTooltips()},
			},
		},
	}
}

var designPalettes = []struct {
	Name  string
	Tone  string
	Steps []int
}{
	{Name: "Zinc", Tone: "zinc", Steps: []int{50, 100, 200, 300, 400, 500, 600, 700, 800, 900, 950}},
	{Name: "Primary (indigo)", Tone: "primary", Steps: []int{50, 100, 200, 300, 400, 500, 600, 700, 800, 900, 950}},
	{Name: "Success", Tone: "success", Steps: []int{50, 100, 200, 300, 400, 500, 600, 700, 800, 900, 950}},
	{Name: "Warning", Tone: "warning", Steps: []int{50, 100, 200, 300, 400, 500, 600, 700, 800, 900, 950}},
	{Name: "Danger", Tone: "danger", Steps: []int{50, 100, 200, 300, 400, 500, 600, 700, 800, 900, 950}},
	{Name: "Info", Tone: "info", Steps: []int{50, 100, 200, 300, 400, 500, 600, 700, 800, 900, 950}},
}

var designSurfaces = []string{"surface-1", "surface-2", "surface-sunken", "surface-overlay"}

var designTextTokens = []string{"text-1", "text-2", "text-muted", "text-inverted"}

var designSpacingSteps = []string{"1", "2", "3", "4", "5", "6", "8", "10", "12"}

var designRadii = []string{"md", "lg"}

var designShadows = []struct{ Name, Var string }{
	{Name: "card", Var: "var(--shadow-card)"},
	{Name: "overlay", Var: "var(--shadow-overlay)"},
	{Name: "modal", Var: "var(--shadow-modal)"},
	{Name: "focus-ring", Var: "var(--focus-ring)"},
	{Name: "focus-ring-danger", Var: "var(--focus-ring-danger)"},
}

var designButtonVariants = []string{"primary", "secondary", "soft", "ghost", "danger", "link"}

var designButtonSizes = []string{"xs", "sm", "md", "lg", "xl"}

var designForceStates = []string{"rest", "hover", "focus", "active"}

var designBadgeTones = []string{"default", "primary", "info", "success", "warning", "danger", "purple", "high"}

var designSemanticTones = []string{"primary", "info", "success", "warning", "danger"}

func designSwatchFG(step int) string {
	if step >= 500 {
		return "white"
	}
	return "black"
}

// designStepperBase pins the "now" anchor so the stepper demo renders the same
// elapsed values every request. Mirrors the TS demo's Date.now() snapshot.
func designStepperBase() time.Time {
	return time.Date(2026, 5, 24, 12, 0, 0, 0, time.UTC)
}
