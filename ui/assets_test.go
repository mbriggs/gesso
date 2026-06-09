package ui

import (
	"encoding/json"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

// The CSS and JS entry points pull in their pieces by relative path, and a
// broken path fails silently in the browser. Resolve every reference against
// the embedded filesystem so a rename or dropped file fails here instead.

var (
	jsImportRe  = regexp.MustCompile(`(?m)^import "ui/(.+)";$`)
	cssImportRe = regexp.MustCompile(`@import url\("\./(.+)"\);`)
)

func entryReferences(t *testing.T, entry string, re *regexp.Regexp) []string {
	t.Helper()

	source, err := fs.ReadFile(embedded, entry)
	if err != nil {
		t.Fatalf("read %s: %v", entry, err)
	}
	matches := re.FindAllStringSubmatch(string(source), -1)
	refs := make([]string, 0, len(matches))
	for _, m := range matches {
		refs = append(refs, m[1])
	}
	return refs
}

func assertResolvable(t *testing.T, refs []string, minimum int) {
	t.Helper()

	if len(refs) < minimum {
		t.Fatalf("found %d references, want at least %d — entry regexp out of sync?", len(refs), minimum)
	}
	for _, ref := range refs {
		content, err := fs.ReadFile(embedded, ref)
		if err != nil {
			t.Errorf("entry references %s but it is not embedded: %v", ref, err)
			continue
		}
		if len(content) == 0 {
			t.Errorf("entry references %s but it is empty", ref)
		}
	}
}

func TestJSEntryImportsResolve(t *testing.T) {
	refs := entryReferences(t, "assets/ui.js", jsImportRe)
	paths := make([]string, 0, len(refs))
	for _, ref := range refs {
		paths = append(paths, "assets/js/"+ref+".js")
	}
	assertResolvable(t, paths, 5)
}

func TestCSSEntryImportsResolve(t *testing.T) {
	refs := entryReferences(t, "assets/ui.css", cssImportRe)
	paths := make([]string, 0, len(refs))
	for _, ref := range refs {
		paths = append(paths, "assets/"+ref)
	}
	assertResolvable(t, paths, 5)
}

func TestImportMapCoversEntryImports(t *testing.T) {
	var parsed struct {
		Imports map[string]string `json:"imports"`
	}
	if err := json.Unmarshal([]byte(importMapJSON), &parsed); err != nil {
		t.Fatalf("import map is not valid JSON: %v", err)
	}

	for _, ref := range entryReferences(t, "assets/ui.js", jsImportRe) {
		url, ok := parsed.Imports["ui/"+ref]
		if !ok {
			t.Errorf("ui.js imports %q but the import map has no entry for it", "ui/"+ref)
			continue
		}
		want := AssetPath("js/" + ref + ".js")
		if url != want {
			t.Errorf("import map maps ui/%s to %q, want %q", ref, url, want)
		}
	}
}

func TestHashedAssetsServeImmutable(t *testing.T) {
	handler := Assets()

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequestWithContext(
		t.Context(), http.MethodGet, "/"+assetHash+"/ui.css", nil,
	))
	if rec.Code != http.StatusOK {
		t.Fatalf("hashed asset = %d, want 200", rec.Code)
	}
	if cc := rec.Header().Get("Cache-Control"); !strings.Contains(cc, "immutable") {
		t.Fatalf("hashed asset Cache-Control = %q, want immutable", cc)
	}

	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequestWithContext(
		t.Context(), http.MethodGet, "/ui.css", nil,
	))
	if rec.Code != http.StatusOK {
		t.Fatalf("bare asset = %d, want 200", rec.Code)
	}
	if cc := rec.Header().Get("Cache-Control"); strings.Contains(cc, "immutable") {
		t.Fatalf("bare asset Cache-Control = %q, want no immutable header", cc)
	}
}
