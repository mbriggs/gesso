// Shell drawer: on narrow viewports the sidebar ([data-ui-shell-sidebar])
// is an off-canvas drawer. [data-ui-shell-toggle] opens/closes it together
// with the scrim ([data-ui-shell-scrim]); scrim click and Escape close.
// Desktop widths pin the sidebar via container query (shell.css) and these
// classes become inert.
function setOpen(shell, open) {
  shell.querySelector("[data-ui-shell-sidebar]")?.classList.toggle("open", open);
  shell.querySelector("[data-ui-shell-scrim]")?.classList.toggle("open", open);
  shell.querySelectorAll("[data-ui-shell-toggle]").forEach((toggle) => {
    toggle.setAttribute("aria-expanded", open ? "true" : "false");
  });
}

document.addEventListener("click", (event) => {
  const toggle = event.target.closest("[data-ui-shell-toggle]");
  if (toggle) {
    const shell = toggle.closest(".shell") || document;
    const sidebar = shell.querySelector("[data-ui-shell-sidebar]");
    setOpen(shell, !(sidebar && sidebar.classList.contains("open")));
    return;
  }
  const scrim = event.target.closest("[data-ui-shell-scrim]");
  if (scrim) setOpen(scrim.closest(".shell") || document, false);
});

document.addEventListener("keydown", (event) => {
  if (event.key !== "Escape") return;
  document.querySelectorAll("[data-ui-shell-sidebar].open").forEach((sidebar) => {
    setOpen(sidebar.closest(".shell") || document, false);
  });
});
