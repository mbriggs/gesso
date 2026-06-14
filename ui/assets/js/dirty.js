// Unsaved-changes guard: forms marked [data-ui-dirty-guard] compare every
// control against its baseline — data-baseline-value when present, else the
// parsed default (defaultValue / defaultChecked / defaultSelected). While
// any control differs the form carries data-ui-dirty="true" (a CSS hook),
// every count change dispatches a bubbling "ui:dirty-change" CustomEvent
// with { form, dirty, count }, and leaving the page warns. Submitting or
// resetting the form stands down. After programmatic edits, dispatch a
// bubbling "change" on the form to force re-evaluation. Controls marked
// [data-ui-dirty-ignore] are never tracked — UI-only inputs (filter
// boxes) that live inside the form but aren't part of its payload.
const skippedTypes = new Set(["button", "submit", "reset", "image", "fieldset", "output"]);
const lastCount = new WeakMap();
const submitting = new WeakSet();

function isTracked(control) {
  if (!control.type || skippedTypes.has(control.type)) return false;
  if (control.hasAttribute("data-ui-dirty-ignore")) return false;
  if (control.type === "hidden" && !control.hasAttribute("data-baseline-value")) return false;
  return true;
}

function controlState(control) {
  if (control.type === "checkbox" || control.type === "radio") {
    return control.checked ? "true" : "false";
  }
  if (control instanceof HTMLSelectElement) {
    return Array.from(control.selectedOptions, (option) => option.value).join("\n");
  }
  return control.value;
}

function controlBaseline(control) {
  const explicit = control.getAttribute("data-baseline-value");
  if (explicit !== null) return explicit;
  if (control.type === "checkbox" || control.type === "radio") {
    return control.defaultChecked ? "true" : "false";
  }
  if (control instanceof HTMLSelectElement) {
    const defaults = Array.from(control.options)
      .filter((option) => option.defaultSelected)
      .map((option) => option.value);
    if (defaults.length === 0 && !control.multiple && control.options.length > 0) {
      return control.options[0].value;
    }
    return defaults.join("\n");
  }
  return control.defaultValue ?? "";
}

function refresh(form) {
  let count = 0;
  for (const control of form.elements) {
    if (!isTracked(control)) continue;
    if (controlState(control) !== controlBaseline(control)) count++;
  }
  const dirty = count > 0;
  if (dirty) form.setAttribute("data-ui-dirty", "true");
  else form.removeAttribute("data-ui-dirty");
  if (lastCount.get(form) === count) return;
  lastCount.set(form, count);
  form.dispatchEvent(
    new CustomEvent("ui:dirty-change", { bubbles: true, detail: { form, dirty, count } }),
  );
}

function guardedForm(target) {
  return target.closest?.("form[data-ui-dirty-guard]");
}

document.addEventListener("input", (event) => {
  const form = guardedForm(event.target);
  if (form) refresh(form);
});

document.addEventListener("change", (event) => {
  const form = guardedForm(event.target);
  if (form) refresh(form);
});

document.addEventListener("reset", (event) => {
  const form = event.target;
  if (!form.matches?.("form[data-ui-dirty-guard]")) return;
  // reset fires before values revert; re-evaluate after they have.
  setTimeout(() => refresh(form), 0);
});

document.addEventListener("submit", (event) => {
  const form = event.target;
  if (!form.matches?.("form[data-ui-dirty-guard]")) return;
  // Another handler (confirm interceptor, swap) may cancel this submit;
  // only stand down once the dispatch settles uncancelled.
  setTimeout(() => {
    if (!event.defaultPrevented) submitting.add(form);
  }, 0);
});

window.addEventListener("beforeunload", (event) => {
  const dirtyForm = Array.from(
    document.querySelectorAll('form[data-ui-dirty-guard][data-ui-dirty="true"]'),
  ).find((form) => !submitting.has(form));
  if (!dirtyForm) return;
  event.preventDefault();
  event.returnValue = "";
});
