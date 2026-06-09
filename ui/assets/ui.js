// Design-system behavior entry point, loaded from the layout as
// <script type="module" src={ ui.AssetPath("ui.js") }>. The bare "ui/*"
// specifiers resolve through the import map rendered by ui.ImportMap()
// (which must precede this script in <head>). Each module registers
// delegated listeners for one component's data-ui-* contract; the contract
// is documented at the top of each file. A page that never renders a
// component pays nothing beyond the no-op listener — modules hold no
// per-page state at load.
import "ui/theme";
import "ui/toggle";
import "ui/segmented";
import "ui/menu";
import "ui/modal";
import "ui/combobox";
import "ui/scrollspy";
