// Swap: forms marked [data-ui-swap] submit via fetch and patch the page in
// place instead of navigating. The contract, kept deliberately small:
//
//   request   the form's own method/action, body from FormData (plus the
//             submitter's name/value), with an "X-UI-Swap: 1" header so
//             handlers can branch fragment-render vs redirect.
//   response  text/html where every TOP-LEVEL element carries an id; each
//             fragment replaces the existing document element with that id
//             (extra fragments without a counterpart are ignored). Error
//             re-renders (e.g. 422) use the same shape.
//   redirect  a redirected response is followed as a full navigation, so
//             handlers that don't fragment-render keep POST→redirect→GET.
//
// After patching, a "ui:swap" CustomEvent fires on document with
// { form, ids } so app code can restore focus or re-derive state. Without
// JS the form submits natively — the server's non-swap path renders the
// full page.
function swapFragments(html, form) {
  const doc = new DOMParser().parseFromString(html, "text/html");
  const ids = [];
  for (const fragment of Array.from(doc.body.children)) {
    if (!fragment.id) continue;
    const target = document.getElementById(fragment.id);
    if (!target) continue;
    target.replaceWith(fragment);
    ids.push(fragment.id);
  }
  document.dispatchEvent(new CustomEvent("ui:swap", { detail: { form, ids } }));
}

document.addEventListener("submit", (event) => {
  if (event.defaultPrevented) return;
  const form = event.target;
  if (!form.matches?.("form[data-ui-swap]")) return;
  event.preventDefault();

  const body = new FormData(form);
  const submitter = event.submitter;
  if (submitter?.name) body.append(submitter.name, submitter.value);

  const method = (form.getAttribute("method") || "get").toUpperCase();
  let url = form.getAttribute("action") || window.location.href;
  const options = {
    method,
    headers: { "X-UI-Swap": "1" },
    credentials: "same-origin",
  };
  if (method === "GET") {
    url = url.split("?")[0] + "?" + new URLSearchParams(body).toString();
  } else {
    options.body = body;
  }

  form.setAttribute("aria-busy", "true");
  fetch(url, options)
    .then(async (response) => {
      if (response.redirected) {
        window.location.assign(response.url);
        return;
      }
      const type = response.headers.get("content-type") || "";
      if (!type.includes("text/html")) throw new Error(`unexpected response: ${response.status}`);
      swapFragments(await response.text(), form);
    })
    .catch(() => {
      // Network failure or a non-fragment response: reload so the server
      // state (and any flash) becomes visible rather than silently stale.
      window.location.reload();
    })
    .finally(() => {
      form.removeAttribute("aria-busy");
    });
});
