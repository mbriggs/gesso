package ui

type TabItem struct {
	Label string
	Href  string
	Count string // optional count bubble after the label
	// Current marks the active tab (aria-current="page"). The server decides
	// which tab is current — there is no client-side panel switching; each
	// tab is a normal navigation.
	Current bool
}

type TabsProps struct {
	Items     []TabItem
	AriaLabel string // default "Tabs"
	Class     string
}
