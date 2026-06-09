# ui

The component package: typed templ components, tone/size class helpers
(`classes.go`), and the embedded assets (`assets/`) served by `Assets()`.
See the repo README for mounting and layout wiring.

```templ
@ui.Button(ui.ButtonProps{Type: "submit", Variant: "primary", Text: "Save"})
@ui.Field(ui.FieldProps{
	ID: "email",
	Label: "Email",
	Control: ui.Input(ui.InputProps{ID: "email", Name: "email"}),
})
```

Conventions:

- One `.templ` file per component area, with its props in a sibling `.go`
  file; generated `*_templ.go` is committed.
- Behavior JS lives one module per component contract in `assets/js/`,
  keyed by `data-ui-*` attributes documented at the top of each file.
- Styles are per-area sheets under `assets/styles/`, imported by `ui.css`
  in layer order. `cmd/cssdead` fails the build if a class loses its last
  reference — every component's states belong in `gallery/` so the demos
  keep them reachable.
