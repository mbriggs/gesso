// Scrollspy: sections marked [data-ui-scrollspy-section][id="..."] map to
// nav links [data-ui-scrollspy-link][href="#..."]. As sections cross into
// view the matching link gets data-active="true" + aria-current="location".
// Also honors the URL hash on load and hashchange.
const sections = Array.from(
  document.querySelectorAll("[data-ui-scrollspy-section][id]"),
);
const links = Array.from(document.querySelectorAll("[data-ui-scrollspy-link]"));

if (sections.length > 0 && links.length > 0) {
  const ids = new Set(sections.map((s) => s.id));

  const activate = (id) => {
    links.forEach((link) => {
      const target = (link.getAttribute("href") || "").replace(/^#/, "");
      const active = target === id;
      if (active) {
        link.setAttribute("data-active", "true");
        link.setAttribute("aria-current", "location");
      } else {
        link.removeAttribute("data-active");
        link.removeAttribute("aria-current");
      }
    });
  };

  const setFromHash = () => {
    const id = window.location.hash.slice(1);
    if (ids.has(id)) activate(id);
  };
  setFromHash();
  window.addEventListener("hashchange", setFromHash);

  if (typeof IntersectionObserver !== "undefined") {
    const observer = new IntersectionObserver(
      (entries) => {
        const visible = entries
          .filter((entry) => entry.isIntersecting)
          .sort((a, b) => a.boundingClientRect.top - b.boundingClientRect.top);
        const next = visible[0]?.target.id;
        if (next) activate(next);
      },
      { rootMargin: "-18% 0px -72% 0px", threshold: [0, 0.2, 1] },
    );
    sections.forEach((section) => observer.observe(section));
  }
}
