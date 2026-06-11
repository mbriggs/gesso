// Global confirm: one dialog (ui.GlobalConfirm, id "ui-confirm") intercepts
// every element that opts in via data-ui-confirm — the attribute value is
// the dialog message. Works on forms (intercepts submit), submit buttons
// (the submitter's attributes win over the form's), and links (intercepts
// click, follows href on accept). Per-element overrides:
//   data-ui-confirm-title         dialog title    (default "Are you sure?")
//   data-ui-confirm-label         confirm button  (default "Confirm")
//   data-ui-confirm-cancel-label  cancel button   (default "Cancel")
//   data-ui-confirm-tone="danger" danger-styled confirm button
// Without the dialog in the DOM (or without JS) nothing blocks — actions
// proceed as plain navigations/submissions.
let pending = null;

function confirmDialog() {
  return document.getElementById("ui-confirm");
}

function openConfirm(source, action) {
  const dialog = confirmDialog();
  if (!dialog || typeof dialog.showModal !== "function") {
    action();
    return;
  }
  dialog.querySelector(".modal-title").textContent =
    source.getAttribute("data-ui-confirm-title") || "Are you sure?";
  dialog.querySelector("[data-ui-confirm-message]").textContent =
    source.getAttribute("data-ui-confirm") || "";
  const accept = dialog.querySelector("[data-ui-confirm-accept]");
  accept.textContent = source.getAttribute("data-ui-confirm-label") || "Confirm";
  const danger = source.getAttribute("data-ui-confirm-tone") === "danger";
  accept.classList.toggle("danger", danger);
  accept.classList.toggle("primary", !danger);
  dialog.querySelector("[data-ui-confirm-cancel]").textContent =
    source.getAttribute("data-ui-confirm-cancel-label") || "Cancel";
  pending = action;
  if (!dialog.open) dialog.showModal();
}

document.addEventListener("submit", (event) => {
  const form = event.target;
  if (!(form instanceof HTMLFormElement)) return;
  if (form.dataset.uiConfirmed) {
    delete form.dataset.uiConfirmed;
    return;
  }
  const submitter = event.submitter;
  const source = submitter?.hasAttribute("data-ui-confirm")
    ? submitter
    : form.hasAttribute("data-ui-confirm")
      ? form
      : null;
  if (!source) return;
  event.preventDefault();
  openConfirm(source, () => {
    form.dataset.uiConfirmed = "true";
    if (submitter) form.requestSubmit(submitter);
    else form.requestSubmit();
  });
});

document.addEventListener("click", (event) => {
  const accept = event.target.closest("[data-ui-confirm-accept]");
  if (accept) {
    const dialog = accept.closest("dialog");
    const action = pending;
    pending = null;
    if (dialog && dialog.open) dialog.close();
    if (action) action();
    return;
  }
  const link = event.target.closest("a[data-ui-confirm]");
  if (!link) return;
  event.preventDefault();
  openConfirm(link, () => {
    window.location.assign(link.href);
  });
});

// Cancel, backdrop click, and ESC all close the dialog without running the
// pending action; "close" doesn't bubble, so listen in the capture phase.
document.addEventListener(
  "close",
  (event) => {
    if (event.target.id === "ui-confirm") pending = null;
  },
  true,
);
