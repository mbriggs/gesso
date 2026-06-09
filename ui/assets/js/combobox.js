// Combobox: filter listbox by input text, keyboard nav (Arrow/Home/End/
// Enter/Escape), commit on click. Hidden input [data-ui-combobox-value]
// mirrors the chosen value so the parent <form> submits like a native
// <select>.
const cbActive = new WeakMap();

function cbOptions(group) {
  return Array.from(group.querySelectorAll("[data-ui-combobox-option]"));
}
function cbVisibleOptions(group) {
  return cbOptions(group).filter((o) => !o.hidden);
}
function cbSetActive(group, index) {
  const visible = cbVisibleOptions(group);
  if (visible.length === 0) {
    cbActive.set(group, -1);
    const input = group.querySelector("[data-ui-combobox-input]");
    if (input) input.removeAttribute("aria-activedescendant");
    return;
  }
  const clamped = Math.max(0, Math.min(index, visible.length - 1));
  cbActive.set(group, clamped);
  visible.forEach((opt, i) => opt.classList.toggle("active", i === clamped));
  const input = group.querySelector("[data-ui-combobox-input]");
  if (input) input.setAttribute("aria-activedescendant", visible[clamped].id);
}
function cbOpen(group, open) {
  const listbox = group.querySelector("[data-ui-combobox-listbox]");
  const input = group.querySelector("[data-ui-combobox-input]");
  if (!listbox || !input) return;
  listbox.classList.toggle("hidden", !open);
  input.setAttribute("aria-expanded", open ? "true" : "false");
  if (!open) {
    cbOptions(group).forEach((o) => o.classList.remove("active"));
    input.removeAttribute("aria-activedescendant");
  }
}
function cbFilter(group, query) {
  const lower = query.toLowerCase();
  let anyVisible = false;
  cbOptions(group).forEach((opt) => {
    const blank = opt.classList.contains("blank");
    const match = blank || (opt.getAttribute("data-label") || "").toLowerCase().includes(lower);
    opt.hidden = !match;
    if (match) anyVisible = true;
  });
  const empty = group.querySelector("[data-ui-combobox-empty]");
  if (empty) empty.hidden = anyVisible;
}
function cbCommit(group, opt) {
  const hidden = group.querySelector("[data-ui-combobox-value]");
  const input = group.querySelector("[data-ui-combobox-input]");
  const value = opt.getAttribute("data-value") || "";
  const label = opt.getAttribute("data-label") || "";
  if (hidden && hidden.value !== value) {
    hidden.value = value;
    hidden.dispatchEvent(new Event("change", { bubbles: true }));
  }
  if (input) input.value = label;
  cbOptions(group).forEach((o) =>
    o.setAttribute("aria-selected", o === opt ? "true" : "false"),
  );
  cbOpen(group, false);
}

document.addEventListener("focusin", (event) => {
  const input = event.target.closest("[data-ui-combobox-input]");
  if (!input) return;
  const group = input.closest("[data-ui-combobox]");
  if (!group) return;
  cbFilter(group, "");
  cbOpen(group, true);
  cbSetActive(group, 0);
});

document.addEventListener("focusout", (event) => {
  const input = event.target.closest("[data-ui-combobox-input]");
  if (!input) return;
  const group = input.closest("[data-ui-combobox]");
  if (!group) return;
  setTimeout(() => {
    // Skip closing if focus moved to an option (mousedown.preventDefault
    // keeps focus on the input, but defensive).
    if (group.contains(document.activeElement)) return;
    cbOpen(group, false);
    const hidden = group.querySelector("[data-ui-combobox-value]");
    const value = hidden ? hidden.value : "";
    const match = cbOptions(group).find((o) => o.getAttribute("data-value") === value);
    input.value = match ? match.getAttribute("data-label") || "" : "";
  }, 0);
});

document.addEventListener("input", (event) => {
  const input = event.target.closest("[data-ui-combobox-input]");
  if (!input) return;
  const group = input.closest("[data-ui-combobox]");
  if (!group) return;
  cbFilter(group, input.value);
  cbOpen(group, true);
  cbSetActive(group, 0);
  const hidden = group.querySelector("[data-ui-combobox-value]");
  if (hidden && hidden.value !== "") {
    hidden.value = "";
    hidden.dispatchEvent(new Event("change", { bubbles: true }));
  }
});

document.addEventListener("keydown", (event) => {
  const input = event.target.closest("[data-ui-combobox-input]");
  if (!input) return;
  const group = input.closest("[data-ui-combobox]");
  if (!group) return;
  const visible = cbVisibleOptions(group);
  const idx = cbActive.get(group) ?? 0;
  if (event.key === "ArrowDown") {
    event.preventDefault();
    cbOpen(group, true);
    cbSetActive(group, Math.min(idx + 1, visible.length - 1));
  } else if (event.key === "ArrowUp") {
    event.preventDefault();
    cbOpen(group, true);
    cbSetActive(group, Math.max(idx - 1, 0));
  } else if (event.key === "Home") {
    event.preventDefault();
    cbSetActive(group, 0);
  } else if (event.key === "End") {
    event.preventDefault();
    cbSetActive(group, visible.length - 1);
  } else if (event.key === "Enter") {
    if (visible.length > 0 && idx >= 0) {
      event.preventDefault();
      cbCommit(group, visible[idx]);
    }
  } else if (event.key === "Escape") {
    event.preventDefault();
    cbOpen(group, false);
  } else if (event.key === "Tab") {
    cbOpen(group, false);
  }
});

document.addEventListener("mousedown", (event) => {
  const opt = event.target.closest("[data-ui-combobox-option]");
  if (!opt) return;
  const group = opt.closest("[data-ui-combobox]");
  if (!group) return;
  event.preventDefault();
  cbCommit(group, opt);
});
