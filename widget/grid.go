package widget

import "gioui.org/layout"

// Grid holds the persistent state for a layout.List that has a
// scrollbar attached.
type Grid struct {
	VScrollbar Scrollbar
	HScrollbar Scrollbar
	layout.Grid
}
