# gesso

A server-rendered design system for Go: [templ](https://github.com/a-h/templ)
components with their CSS and JS embedded in the module — the prepared
surface a host app paints its own coat over. No Node, no Tailwind, no asset
pipeline: import the module, mount one route, and the whole system ships
inside your binary.

The only dependency is `a-h/templ`.

## Usage

```go
import "github.com/mbriggs/gesso/ui"

// Serve the embedded assets (any router; echo shown):
e.GET("/ui/*", echo.WrapHandler(http.StripPrefix("/ui/", ui.Assets())))
```

In your layout `<head>`:

```templ
<link rel="stylesheet" href={ ui.AssetPath("ui.css") }/>
<script>document.documentElement.classList.add("js");</script>
@ui.ImportMap()
<script type="module" src={ ui.AssetPath("ui.js") }></script>
```

Then compose pages from the components:

```templ
@ui.Button(ui.ButtonProps{Type: "submit", Text: "Save"})
@ui.Field(ui.FieldProps{ID: "email", Label: "Email", Control: ui.Input(...)})
```

`ui.AssetPath` returns content-hashed URLs (`/ui/<hash>/ui.css`) served with
an immutable cache header; the hash changes when any embedded asset changes.
`ui.ImportMap()` resolves the bare `ui/*` module specifiers `ui.js` imports —
one JS file per component behavior under `ui/assets/js/`, each documenting
its `data-ui-*` contract.

## Gallery

`gallery.Page` renders every component and state with a scrollspy section
nav. Mount it behind a dev-only route in your own layout:

```go
page := gallery.PageData{Groups: gallery.Sections()}
// render gallery.Page(page) inside your layout,
// with gallery.HeadExtras() in <head>
```

## Conventions

- Styles are layered: `@layer theme, base, components, app`. The module owns
  the first three; the host's own stylesheet is the `app` layer.
- `cmd/cssdead` (run by `bin/check`) fails on CSS classes no component or
  gallery demo references — deleting a component means deleting its styles.
  Render-time-constructed families (`align-*`, `gap-*`, `*-spacing`) are
  allowlisted in its main.go; `/* cssdead:allow */` exempts a single rule.
- Generated `*_templ.go` is committed; `bin/check` fails on drift. Hosts
  never need the templ CLI unless they edit components.

## Development

```sh
mise install   # go, templ, gofumpt, golangci-lint, shellcheck
bin/check      # the full gate
```

When developing against a host app, point the host at your checkout:

```
// host go.mod
replace github.com/mbriggs/gesso => ../gesso
```
