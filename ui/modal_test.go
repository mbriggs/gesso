package ui_test

import (
	"strings"
	"testing"

	"github.com/mbriggs/gesso/ui"
)

func TestGlobalConfirmCarriesInterceptorHooks(t *testing.T) {
	got := renderComponent(t, ui.GlobalConfirm())
	for _, want := range []string{
		`id="ui-confirm"`,
		`data-ui-modal`,
		`id="ui-confirm-title"`,
		`>Are you sure?</h3>`,
		`<span data-ui-confirm-message></span>`,
		`data-ui-confirm-accept`,
		`data-ui-confirm-cancel`,
		`data-ui-modal-close`,
		`>Confirm</button>`,
		`>Cancel</button>`,
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("global confirm missing %q:\n%s", want, got)
		}
	}
	if strings.Contains(got, " open") {
		t.Fatalf("dialog must not render the open attribute:\n%s", got)
	}
}
