package ui_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/a-h/templ"
	"github.com/mbriggs/gesso/ui"
)

func TestTemplComponentsRenderButton(t *testing.T) {
	got := renderComponent(t, ui.Button(ui.ButtonProps{Variant: "secondary", Size: "sm", Text: "Save"}))
	for _, want := range []string{`<main>`, `class="button secondary sm shadow"`, `Save`} {
		if !strings.Contains(got, want) {
			t.Fatalf("rendered page = %q, want %q", got, want)
		}
	}
}

func TestFeedbackComponentsRenderTypedClasses(t *testing.T) {
	got := renderComponent(t, ui.Components(
		ui.Alert(ui.AlertProps{Tone: ui.ToneDanger, Class: "extra", Text: "Bad input"}),
		ui.Badge(ui.BadgeProps{Tone: ui.TonePrimary, Pill: true, Text: "Ready"}),
		ui.Progress(ui.ProgressProps{Value: 2, Total: 4, Tone: ui.ToneSuccess, Size: "xs", Inline: true, Label: "Half done"}),
	))
	for _, want := range []string{
		`role="alert" class="alert danger extra"`,
		`class="badge primary pill"`,
		`class="progress xs inline"`,
		`class="progress-bar success"`,
		`style="--progress-value: 50%;"`,
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("rendered page missing %q:\n%s", want, got)
		}
	}
}

func TestFormComponentsCarryAccessibleState(t *testing.T) {
	got := renderComponent(t, ui.Components(
		ui.Field(ui.FieldProps{
			ID:    "email",
			Label: "Email",
			Hint:  "Use work email.",
			Error: "Required.",
			Control: ui.Input(ui.InputProps{
				ID:    "email",
				Name:  "email",
				Hint:  "Use work email.",
				Error: "Required.",
			}),
		}),
		ui.Segmented(ui.SegmentedProps{
			ID:    "mode",
			Name:  "mode",
			Value: "manual",
			Label: "Mode",
			Hint:  "Choose one.",
			Choices: []ui.SegmentedChoice{
				{Label: "Auto", Value: "auto"},
				{Label: "Manual", Value: "manual"},
			},
		}),
		ui.Toggle(ui.ToggleProps{
			ID:      "enabled",
			Name:    "enabled",
			Label:   "Enabled",
			Hint:    "Sync later.",
			Checked: true,
		}),
	))
	for _, want := range []string{
		`id="email-hint"`,
		`id="email-error"`,
		`aria-describedby="email-hint email-error"`,
		`aria-invalid="true"`,
		`type="hidden" name="mode" value="manual"`,
		`data-ui-segmented`,
		`aria-labelledby="mode-label"`,
		`tabindex="0"`,
		`data-ui-toggle-input`,
		`tabindex="-1"`,
		`aria-labelledby="enabled-label"`,
		`aria-describedby="enabled-hint"`,
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("rendered page missing %q:\n%s", want, got)
		}
	}
}

func TestAssetsServeStylesheet(t *testing.T) {
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/ui.css", nil)
	w := httptest.NewRecorder()

	ui.Assets().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("code = %d, want 200", w.Code)
	}
	body, err := io.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if !strings.Contains(string(body), `@import url("./styles/tokens.css")`) {
		t.Fatalf("stylesheet body = %q", body)
	}
}

func renderComponent(t *testing.T, component templ.Component) string {
	t.Helper()
	var b bytes.Buffer
	if _, err := b.WriteString("<main>"); err != nil {
		t.Fatalf("write main: %v", err)
	}
	if err := component.Render(context.Background(), &b); err != nil {
		t.Fatalf("render component: %v", err)
	}
	if _, err := b.WriteString("</main>"); err != nil {
		t.Fatalf("write main: %v", err)
	}
	return b.String()
}
