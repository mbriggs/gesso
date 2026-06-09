// Theme: applies the stored light/dark preference on load and flips it via
// [data-ui-theme-toggle]. The dark palette lives in styles/components/dark.css
// behind the .is-dark class on <html>.
const root = document.documentElement;

try {
  root.classList.toggle("is-dark", localStorage.getItem("theme") === "dark");
} catch {
  // Leave the server-rendered theme alone when storage is unavailable.
}

document.addEventListener("click", (event) => {
  const button = event.target.closest("[data-ui-theme-toggle]");
  if (!button) return;

  const next = !root.classList.contains("is-dark");
  root.classList.toggle("is-dark", next);
  try {
    localStorage.setItem("theme", next ? "dark" : "light");
  } catch {
    // The in-memory class change is still useful in restricted contexts.
  }
});
