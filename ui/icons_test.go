package ui_test

import (
	"strings"
	"testing"

	"github.com/mbriggs/gesso/ui"
)

func TestEveryIconNameRendersAGlyph(t *testing.T) {
	names := ui.IconNames()
	if len(names) == 0 {
		t.Fatal("IconNames returned no icons")
	}
	for _, name := range names {
		got := renderComponent(t, ui.Icon(name, ""))
		if !strings.Contains(got, "<path") {
			t.Fatalf("icon %q renders no paths:\n%s", name, got)
		}
		if !strings.Contains(got, `class="icon"`) {
			t.Fatalf("icon %q missing .icon class:\n%s", name, got)
		}
		if !strings.Contains(got, `aria-hidden="true"`) {
			t.Fatalf("icon %q not aria-hidden:\n%s", name, got)
		}
	}
}

func TestIconVariantsAndUnknownNames(t *testing.T) {
	solid := renderComponent(t, ui.Icon(ui.IconStarSolid, "extra"))
	if !strings.Contains(solid, `fill="currentColor"`) || strings.Contains(solid, `stroke="currentColor"`) {
		t.Fatalf("solid icon should fill, not stroke:\n%s", solid)
	}
	if !strings.Contains(solid, `class="icon extra"`) {
		t.Fatalf("solid icon missing extra class:\n%s", solid)
	}

	outline := renderComponent(t, ui.Icon(ui.IconStar, ""))
	if !strings.Contains(outline, `fill="none"`) || !strings.Contains(outline, `stroke="currentColor"`) {
		t.Fatalf("outline icon should stroke:\n%s", outline)
	}

	unknown := renderComponent(t, ui.Icon(ui.IconName("nope"), ""))
	if strings.Contains(unknown, "<path") {
		t.Fatalf("unknown icon should render an empty glyph:\n%s", unknown)
	}
}
