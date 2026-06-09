package ui

import "github.com/a-h/templ"

type ModalSize string

const (
	ModalSizeMd   ModalSize = "md"
	ModalSizeWide ModalSize = "wide"
	ModalSizeFull ModalSize = "full"
)

type ModalProps struct {
	ID      string
	TitleID string
	Title   string
	Size    ModalSize
	// AutoOpen attaches a marker that the ui.js handler picks up on
	// DOMContentLoaded and calls showModal() so the dialog gets native
	// modal semantics (backdrop, focus trap, ESC). Setting the HTML
	// `open` attribute does NOT produce modal behavior — only
	// showModal() does — so we never emit it from markup.
	AutoOpen        bool
	NoBackdropClose bool
	HeaderActions   templ.Component
	Content         templ.Component
}

type ModalBodyProps struct {
	Class   string
	Content templ.Component
}

type ConfirmTone string

const (
	ConfirmToneDefault ConfirmTone = "default"
	ConfirmToneDanger  ConfirmTone = "danger"
)

type ConfirmDialogProps struct {
	ID                 string
	AutoOpen           bool
	Tone               ConfirmTone
	Title              string
	Description        string
	DescriptionContent templ.Component
	ConfirmLabel       string
	CancelLabel        string
	Busy               bool
	// When set, the confirm button submits the named form (use with a form
	// elsewhere on the page). Otherwise the dialog assumes the caller wires
	// behavior via ConfirmAttrs.
	ConfirmForm  string
	ConfirmName  string
	ConfirmValue string
	ConfirmAttrs templ.Attributes
	// CancelAttrs overrides the default close-the-modal wiring. Leave nil
	// to use the built-in `data-ui-modal-close` handler.
	CancelAttrs templ.Attributes
}

func modalSizeClass(size ModalSize) string {
	if size == "" || size == ModalSizeMd {
		return ""
	}
	return string(size)
}

func confirmDialogConfirmType(p ConfirmDialogProps) string {
	if p.ConfirmForm != "" {
		return "submit"
	}
	return "button"
}

func confirmDialogConfirmVariant(p ConfirmDialogProps) string {
	if p.Tone == ConfirmToneDanger {
		return "danger"
	}
	return "primary"
}

func confirmDialogCancelAttrs(p ConfirmDialogProps) templ.Attributes {
	if p.CancelAttrs != nil {
		return p.CancelAttrs
	}
	return templ.Attributes{"data-ui-modal-close": ""}
}
