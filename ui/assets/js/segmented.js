// Segmented control: radio-style option row. [data-ui-segmented-option]
// buttons inside [data-ui-segmented] update aria-checked/tabindex and mirror
// the choice into the hidden [data-ui-segmented-input] so the parent form
// sees a normal field. Arrow/Home/End keys move the selection.
const selectSegmented = (group, option) => {
  const input = group.querySelector("[data-ui-segmented-input]");
  const selected = option.getAttribute("aria-checked") === "true";
  if (selected && (!input || input.value === option.value)) return;

  group.querySelectorAll("[data-ui-segmented-option]").forEach((candidate) => {
    const active = candidate === option;
    candidate.setAttribute("aria-checked", active ? "true" : "false");
    candidate.tabIndex = active ? 0 : -1;
  });
  if (input) {
    input.value = option.value;
    input.dispatchEvent(new Event("change", { bubbles: true }));
  }
};

document.addEventListener("click", (event) => {
  const option = event.target.closest("[data-ui-segmented-option]");
  if (!option) return;

  const group = option.closest("[data-ui-segmented]");
  if (!group) return;
  selectSegmented(group, option);
});

document.addEventListener("keydown", (event) => {
  const current = event.target.closest("[data-ui-segmented-option]");
  if (!current) return;
  const group = current.closest("[data-ui-segmented]");
  if (!group) return;

  const options = Array.from(group.querySelectorAll("[data-ui-segmented-option]"));
  if (options.length === 0) return;

  const index = options.indexOf(current);
  let nextIndex = -1;
  if (event.key === "ArrowRight" || event.key === "ArrowDown") {
    nextIndex = index >= options.length - 1 ? 0 : index + 1;
  } else if (event.key === "ArrowLeft" || event.key === "ArrowUp") {
    nextIndex = index <= 0 ? options.length - 1 : index - 1;
  } else if (event.key === "Home") {
    nextIndex = 0;
  } else if (event.key === "End") {
    nextIndex = options.length - 1;
  }
  if (nextIndex < 0) return;

  event.preventDefault();
  const next = options[nextIndex];
  selectSegmented(group, next);
  next.focus();
});
