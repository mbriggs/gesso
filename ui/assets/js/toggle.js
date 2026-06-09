// Toggle switch: [data-ui-toggle-switch] is the styled button face over a
// real checkbox [data-ui-toggle-input] inside [data-ui-toggle]. Clicking the
// face flips the checkbox (so forms submit normally); checkbox changes sync
// aria-checked back onto the face.
document.addEventListener("click", (event) => {
  const button = event.target.closest("[data-ui-toggle-switch]");
  if (!button || button.disabled) return;

  const toggle = button.closest("[data-ui-toggle]");
  const input = toggle?.querySelector('input[type="checkbox"]');
  if (!input || input.disabled) return;

  input.checked = !input.checked;
  input.dispatchEvent(new Event("change", { bubbles: true }));
});

document.addEventListener("change", (event) => {
  const input = event.target.closest("[data-ui-toggle-input]");
  if (!input) return;

  const button = input.closest("[data-ui-toggle]")?.querySelector("[data-ui-toggle-switch]");
  if (!button) return;
  button.setAttribute("aria-checked", input.checked ? "true" : "false");
});
