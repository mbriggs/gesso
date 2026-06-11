package ui

import "slices"

// IconName selects an icon from the built-in set. The set covers the
// product surfaces this design system serves (nav glyphs, status marks,
// chrome affordances); icons render as 24x24 stroke glyphs unless the
// shape is dot-based (grip, kebab) or has a solid variant (star).
type IconName string

const (
	IconPlus            IconName = "plus"
	IconX               IconName = "x"
	IconCheck           IconName = "check"
	IconCheckCircle     IconName = "check-circle"
	IconXCircle         IconName = "x-circle"
	IconInfoCircle      IconName = "info-circle"
	IconWarningTriangle IconName = "warning-triangle"
	IconChevronUp       IconName = "chevron-up"
	IconChevronDown     IconName = "chevron-down"
	IconChevronLeft     IconName = "chevron-left"
	IconChevronRight    IconName = "chevron-right"
	IconMenu            IconName = "menu"
	IconSearch          IconName = "search"
	IconClock           IconName = "clock"
	IconStar            IconName = "star"
	IconStarSolid       IconName = "star-solid"
	IconFolder          IconName = "folder"
	IconBox             IconName = "box"
	IconStack           IconName = "stack"
	IconWrench          IconName = "wrench"
	IconInbox           IconName = "inbox"
	IconImage           IconName = "image"
	IconMoon            IconName = "moon"
	IconSun             IconName = "sun"
	IconSwatch          IconName = "swatch"
	IconGrip            IconName = "grip"
	IconKebab           IconName = "kebab"
)

type iconShape struct {
	Solid bool
	Paths []string
}

var icons = map[IconName]iconShape{
	IconPlus: {Paths: []string{"M12 5v14", "M5 12h14"}},
	IconX:    {Paths: []string{"M18 6 6 18", "m6 6 12 12"}},
	IconCheck: {Paths: []string{
		"m20 6-11 11-5-5",
	}},
	IconCheckCircle: {Paths: []string{
		"M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z",
		"m8.5 12.5 2.5 2.5 5.5-6",
	}},
	IconXCircle: {Paths: []string{
		"M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z",
		"m15 9-6 6",
		"m9 9 6 6",
	}},
	IconInfoCircle: {Paths: []string{
		"M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z",
		"M12 16v-4.5",
		"M12 8h.01",
	}},
	IconWarningTriangle: {Paths: []string{
		"M10.29 3.86 1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0Z",
		"M12 9v4",
		"M12 17h.01",
	}},
	IconChevronUp:    {Paths: []string{"m18 15-6-6-6 6"}},
	IconChevronDown:  {Paths: []string{"m6 9 6 6 6-6"}},
	IconChevronLeft:  {Paths: []string{"m15 18-6-6 6-6"}},
	IconChevronRight: {Paths: []string{"m9 18 6-6-6-6"}},
	IconMenu:         {Paths: []string{"M4 6h16", "M4 12h16", "M4 18h16"}},
	IconSearch: {Paths: []string{
		"M19 11a8 8 0 1 1-16 0 8 8 0 0 1 16 0Z",
		"m21 21-4.3-4.3",
	}},
	IconClock: {Paths: []string{
		"M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z",
		"M12 7v5l3 1.5",
	}},
	IconStar: {Paths: []string{
		"m12 3 2.65 5.62 6.1.78-4.49 4.28 1.16 6.05L12 16.77l-5.42 2.96 1.16-6.05-4.49-4.28 6.1-.78L12 3Z",
	}},
	IconStarSolid: {Solid: true, Paths: []string{
		"m12 3 2.65 5.62 6.1.78-4.49 4.28 1.16 6.05L12 16.77l-5.42 2.96 1.16-6.05-4.49-4.28 6.1-.78L12 3Z",
	}},
	IconFolder: {Paths: []string{
		"M3 7a2 2 0 0 1 2-2h4.2a2 2 0 0 1 1.4.6L12 7h7a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V7Z",
	}},
	IconBox: {Paths: []string{
		"M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16Z",
		"m3.3 7 8.7 5 8.7-5",
		"M12 22V12",
	}},
	IconStack: {Paths: []string{
		"M5 10h14a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2Z",
		"M6 6.5h12",
		"M8 3h8",
	}},
	IconWrench: {Paths: []string{
		"M14.7 6.3a1 1 0 0 0 0 1.4l1.6 1.6a1 1 0 0 0 1.4 0l3.77-3.77a6 6 0 0 1-7.94 7.94l-6.91 6.91a2.12 2.12 0 0 1-3-3l6.91-6.91a6 6 0 0 1 7.94-7.94l-3.76 3.76Z",
	}},
	IconInbox: {Paths: []string{
		"M22 12h-6l-2 3h-4l-2-3H2",
		"M5.45 5.11 2 12v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-6l-3.45-6.89A2 2 0 0 0 16.76 4H7.24a2 2 0 0 0-1.79 1.11Z",
	}},
	IconImage: {Paths: []string{
		"M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2Z",
		"M11 9a2 2 0 1 1-4 0 2 2 0 0 1 4 0Z",
		"m21 15-3.08-3.08a2 2 0 0 0-2.83 0L6 21",
	}},
	IconMoon: {Paths: []string{
		"M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z",
	}},
	IconSun: {Paths: []string{
		"M16 12a4 4 0 1 1-8 0 4 4 0 0 1 8 0Z",
		"M12 2v2",
		"M12 20v2",
		"m4.93 4.93 1.41 1.41",
		"m17.66 17.66 1.41 1.41",
		"M2 12h2",
		"M20 12h2",
		"m6.34 17.66-1.41 1.41",
		"m19.07 4.93-1.41 1.41",
	}},
	IconSwatch: {Paths: []string{
		"M11 17a4 4 0 0 1-8 0V5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2Z",
		"M16.7 13H19a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H7",
		"m11 8 2.3-2.3a2.4 2.4 0 0 1 3.4 0l2.6 2.6a2.4 2.4 0 0 1 0 3.4l-7.6 7.5",
		"M7 17h.01",
	}},
	IconGrip: {Solid: true, Paths: []string{
		"M10.25 5a1.25 1.25 0 1 1-2.5 0 1.25 1.25 0 0 1 2.5 0Z",
		"M16.25 5a1.25 1.25 0 1 1-2.5 0 1.25 1.25 0 0 1 2.5 0Z",
		"M10.25 12a1.25 1.25 0 1 1-2.5 0 1.25 1.25 0 0 1 2.5 0Z",
		"M16.25 12a1.25 1.25 0 1 1-2.5 0 1.25 1.25 0 0 1 2.5 0Z",
		"M10.25 19a1.25 1.25 0 1 1-2.5 0 1.25 1.25 0 0 1 2.5 0Z",
		"M16.25 19a1.25 1.25 0 1 1-2.5 0 1.25 1.25 0 0 1 2.5 0Z",
	}},
	IconKebab: {Solid: true, Paths: []string{
		"M13.5 5a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0Z",
		"M13.5 12a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0Z",
		"M13.5 19a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0Z",
	}},
}

// IconNames returns every registered icon name, sorted; the gallery renders
// the full set from this so the build proves each name has a shape.
func IconNames() []IconName {
	names := make([]IconName, 0, len(icons))
	for name := range icons {
		names = append(names, name)
	}
	slices.Sort(names)
	return names
}

// iconShapeFor looks up a shape; unknown names render an empty glyph rather
// than panicking, matching the tone-coercion posture elsewhere.
func iconShapeFor(name IconName) iconShape {
	return icons[name]
}
