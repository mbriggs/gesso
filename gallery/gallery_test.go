package gallery_test

import (
	"context"
	"strings"
	"testing"

	"github.com/mbriggs/gesso/gallery"
)

func TestPageRendersEverySectionWithScrollspyWiring(t *testing.T) {
	page := gallery.PageData{Groups: gallery.Sections()}

	var b strings.Builder
	if err := gallery.Page(page).Render(context.Background(), &b); err != nil {
		t.Fatalf("render gallery page: %v", err)
	}
	html := b.String()

	for _, want := range []string{
		"Design system",
		"design-shell",
		"data-ui-scrollspy-link",
	} {
		if !strings.Contains(html, want) {
			t.Fatalf("gallery page missing %q", want)
		}
	}
	for _, group := range page.Groups {
		for _, section := range group.Sections {
			if !strings.Contains(html, `id="`+section.ID+`"`) {
				t.Fatalf("gallery page missing section %q", section.ID)
			}
			if !strings.Contains(html, `href="#`+section.ID+`"`) {
				t.Fatalf("sidebar missing link to %q", section.ID)
			}
		}
	}
}
