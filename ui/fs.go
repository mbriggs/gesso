// Package ui provides templ components, class helpers, and static assets for
// the shared server-rendered design system.
package ui

import (
	"context"
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"path"
	"sort"
	"strings"

	"github.com/a-h/templ"
)

//go:embed assets
var embedded embed.FS

// assetHash fingerprints every embedded asset. It versions asset URLs via
// AssetPath, so a deploy with changed assets gets fresh URLs while unchanged
// ones stay cached forever (the hashed handler serves immutable).
var (
	assets        fs.FS
	assetHash     string
	importMapJSON string
)

func init() {
	sub, err := fs.Sub(embedded, "assets")
	if err != nil {
		panic(err)
	}
	assets = sub
	assetHash = hashAssets(sub)
	importMapJSON = buildImportMap(sub)
}

func hashAssets(root fs.FS) string {
	var paths []string
	err := fs.WalkDir(root, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			paths = append(paths, p)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	sort.Strings(paths)

	h := sha256.New()
	for _, p := range paths {
		content, err := fs.ReadFile(root, p)
		if err != nil {
			panic(err)
		}
		// hash.Hash writes never fail.
		_, _ = fmt.Fprintf(h, "%s\x00", p)
		_, _ = h.Write(content)
	}
	return hex.EncodeToString(h.Sum(nil))[:12]
}

// buildImportMap maps every behavior module to a bare "ui/<name>" specifier,
// so ui.js (and any app module) can `import "ui/combobox"` without knowing
// the hashed URL. Rendered into the layout head by ImportMap.
func buildImportMap(root fs.FS) string {
	entries, err := fs.ReadDir(root, "js")
	if err != nil {
		panic(err)
	}
	imports := make(map[string]string, len(entries))
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".js") {
			continue
		}
		imports["ui/"+strings.TrimSuffix(name, ".js")] = AssetPath("js/" + name)
	}
	out, err := json.Marshal(map[string]any{"imports": imports})
	if err != nil {
		panic(err)
	}
	return string(out)
}

// AssetPath returns the content-hashed URL for an embedded asset, e.g.
// AssetPath("ui.css") → "/ui/<hash>/ui.css". Hashed URLs are served with
// an immutable cache header; the hash changes whenever any asset changes.
func AssetPath(rel string) string {
	return "/ui/" + assetHash + "/" + path.Clean(rel)
}

// ImportMap renders the <script type="importmap"> that resolves bare
// "ui/*" module specifiers to hashed URLs. It must appear in <head> before
// the module script that loads ui.js.
func ImportMap() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		if _, err := fmt.Fprintf(w, `<script type="importmap">%s</script>`, importMapJSON); err != nil {
			return fmt.Errorf("rendering import map: %w", err)
		}
		return nil
	})
}

// Assets returns an HTTP handler for the embedded UI assets. Requests
// prefixed with the current asset hash (what AssetPath emits) are served
// immutable-cached; bare paths still work and use default caching.
func Assets() http.Handler {
	files := http.FileServerFS(assets)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rel := strings.TrimPrefix(r.URL.Path, "/")
		if rest, ok := strings.CutPrefix(rel, assetHash+"/"); ok {
			r = r.Clone(r.Context())
			r.URL.Path = "/" + rest
			w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		}
		files.ServeHTTP(w, r)
	})
}
