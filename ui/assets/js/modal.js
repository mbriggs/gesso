// Modal: open via [data-ui-modal-open="targetId"], close via
// [data-ui-modal-close] inside the dialog or backdrop click. Built on
// native <dialog> so focus trap and ESC come from the browser.
function openModalDialog(dialog) {
  if (!dialog || typeof dialog.showModal !== "function") return;
  if (dialog.open) return;
  try {
    dialog.showModal();
  } catch {
    // showModal throws when the dialog is detached or already open in
    // a non-modal state; swallow rather than break the whole page.
  }
}

// Auto-open dialogs marked by the server. Setting the HTML `open`
// attribute would render a non-modal dialog (no backdrop, no focus
// trap, no top layer), so the server emits a data attribute instead
// and we promote it via showModal() once the document is parsed.
function activateAutoOpenModals(root) {
  (root || document)
    .querySelectorAll("dialog[data-ui-modal][data-ui-modal-auto-open]")
    .forEach(openModalDialog);
}
if (document.readyState === "loading") {
  document.addEventListener("DOMContentLoaded", () => activateAutoOpenModals());
} else {
  activateAutoOpenModals();
}

document.addEventListener("click", (event) => {
  const opener = event.target.closest("[data-ui-modal-open]");
  if (opener) {
    const id = opener.getAttribute("data-ui-modal-open");
    const dialog = id && document.getElementById(id);
    if (dialog && typeof dialog.showModal === "function" && !dialog.open) {
      event.preventDefault();
      openModalDialog(dialog);
    }
    return;
  }
  const closer = event.target.closest("[data-ui-modal-close]");
  if (closer) {
    const dialog = closer.closest("dialog[data-ui-modal]");
    if (dialog && dialog.open) {
      event.preventDefault();
      dialog.close();
    }
  }
});

document.addEventListener("mousedown", (event) => {
  const dialog = event.target.closest("dialog[data-ui-modal]");
  if (!dialog || !dialog.open) return;
  if (dialog.hasAttribute("data-ui-modal-no-backdrop")) return;
  // Click landed inside the panel: ignore. Otherwise it's a backdrop click.
  if (event.target.closest(".modal-panel")) return;
  dialog.close();
});
