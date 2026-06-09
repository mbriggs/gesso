// Menu keyboard nav: while a menu-panel is the active focus context,
// ArrowUp/Down move between [role="menuitem"], Home/End jump to ends,
// and Escape closes the surrounding <details>. Clicking outside closes
// the menu (native <details> handles its own toggle).
function menuItems(panel) {
  return Array.from(panel.querySelectorAll('[role="menuitem"]:not([disabled])'));
}
function menuRoot(el) {
  const details = el.closest("details.menu-root");
  return details && details.open ? details : null;
}

document.addEventListener("keydown", (event) => {
  const target = event.target;
  if (!target || !target.closest) return;
  const insideItem = target.closest('[role="menuitem"]');
  const insidePanel = target.closest(".menu-panel");
  // Match the trigger when focus sits on the <summary> itself or on a
  // direct child (templ may render the trigger as <summary>…inner…</summary>).
  const summary =
    target.matches?.("details.menu-root > summary") ||
    target.parentElement?.matches?.("details.menu-root > summary")
      ? target
      : null;
  const root = menuRoot(target);
  if (!root) return;
  const panel = root.querySelector(".menu-panel");
  if (!panel) return;
  if (event.key === "Escape") {
    event.preventDefault();
    root.open = false;
    const trigger = root.querySelector(":scope > summary");
    if (trigger && typeof trigger.focus === "function") trigger.focus();
    return;
  }
  if (event.key === "ArrowDown" || event.key === "ArrowUp") {
    const items = menuItems(panel);
    if (items.length === 0) return;
    event.preventDefault();
    if (summary || (!insideItem && insidePanel)) {
      items[event.key === "ArrowDown" ? 0 : items.length - 1].focus();
      return;
    }
    if (!insideItem) return;
    const idx = items.indexOf(insideItem);
    const next =
      event.key === "ArrowDown"
        ? (idx + 1) % items.length
        : (idx - 1 + items.length) % items.length;
    items[next].focus();
  } else if (event.key === "Home") {
    const items = menuItems(panel);
    if (items.length === 0) return;
    event.preventDefault();
    items[0].focus();
  } else if (event.key === "End") {
    const items = menuItems(panel);
    if (items.length === 0) return;
    event.preventDefault();
    items[items.length - 1].focus();
  } else if (event.key === "Tab") {
    // Let focus leave naturally, but close the menu so a stale panel isn't
    // left visible behind subsequent focus.
    root.open = false;
  }
});

document.addEventListener("click", (event) => {
  document.querySelectorAll("details.menu-root[open]").forEach((root) => {
    if (root.contains(event.target)) return;
    root.open = false;
  });
});
