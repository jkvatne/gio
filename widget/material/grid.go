// SPDX-License-Identifier: Unlicense OR MIT

package material

import (
	"image"
	"math"

	"gioui.org/f32"
	"gioui.org/op"

	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/unit"
	"gioui.org/widget"
)

// TableStyle is the persistent state for a table with heading and grid.
type TableStyle GridStyle

// Table makes a grid with its persistent state.
func Table(th *Theme, state *widget.Grid) TableStyle {
	return TableStyle{
		State:           state,
		VScrollbarStyle: Scrollbar(th, &state.VScrollbar),
		HScrollbarStyle: Scrollbar(th, &state.HScrollbar),
	}
}

// GridStyle is the persistent state for the grid.
type GridStyle struct {
	State           *widget.Grid
	VScrollbarStyle ScrollbarStyle
	HScrollbarStyle ScrollbarStyle
	AnchorStrategy
}

// Grid makes a grid with its persistent state.
func Grid(th *Theme, state *widget.Grid) GridStyle {
	return GridStyle{
		State:           state,
		VScrollbarStyle: Scrollbar(th, &state.VScrollbar),
		HScrollbarStyle: Scrollbar(th, &state.HScrollbar),
	}
}

// constrain a value to be between min and max (inclusive).
func constrain(value *int, min int, max int) {
	if *value < min {
		*value = min
	}
	if *value > max {
		*value = max
	}
}

// WeightedWidths calculates widths of columns and total width of grid.
// Input is weights[], which are either in Device independent pixels, or
// fractions of actual grid size when less than 1.0. Mixed fractions and
// fixed width columns are allowed. Actual grid size comes from gtx.Max.X.
func WeightedWidths(gtx layout.Context, weights []float32) []unit.Value {
	fracSum := float32(0.0)
	fixWidth := 0
	widths := make([]unit.Value, len(weights))
	for _, w := range weights {
		if w <= 1.0 {
			fracSum += w
		} else {
			fixWidth += gtx.Px(unit.Dp(w))
		}
	}
	scale := float32(gtx.Constraints.Max.X-fixWidth) / fracSum
	sum := 0
	for i := range weights {
		if weights[i] != 0 {
			if weights[i] <= 1.0 {
				widths[i] = unit.Px(float32(math.Round(float64(weights[i] * scale))))
			} else {
				widths[i] = unit.Dp(weights[i])
			}
		} else {
			widths[i] = unit.Px(float32(gtx.Constraints.Max.X / len(weights)))
		}
		sum += gtx.Px(widths[i])
	}
	// Make sure the sum is equal to Max.X
	widths[len(widths)-1] = unit.Add(gtx.Metric, widths[len(widths)-1], unit.Px(float32(gtx.Constraints.Max.X-sum)))
	return widths
}

func drawHeading(gtx layout.Context, headingFunc layout.ListElement, rowHeight int, colWidths []int, HorPos int) int {
	headingHeight := 0
	if headingFunc != nil {
		headingHeight = rowHeight
		c := image.Pt(gtx.Constraints.Max.X, rowHeight)
		cl := clip.Rect{Max: c}.Push(gtx.Ops)
		firstCol, lastCol, colOffset := layout.FindColStart(HorPos, gtx.Constraints.Max.X, colWidths)
		if headingFunc != nil {
			xPos := colOffset
			for col := firstCol; col <= lastCol; col++ {
				c := gtx
				c.Constraints.Max.X = colWidths[col]
				c.Constraints.Min.X = colWidths[col]
				c.Constraints.Max.Y = rowHeight
				c.Constraints.Min.Y = rowHeight
				trans := op.Offset(f32.Pt(float32(xPos), 0.0)).Push(gtx.Ops)
				clp := clip.Rect{Max: c.Constraints.Max}.Push(gtx.Ops)
				d := headingFunc(c, col)
				clp.Pop()
				trans.Pop()
				xPos += d.Size.X
			}
		}
		cl.Pop()
	}
	return headingHeight
}

// Layout will draw a table with a heading, using fixed column widths and row height.
func (t TableStyle) Layout(gtx layout.Context, rowCount int, rowHeightValue unit.Value,
	colWidths []unit.Value, cellFunc layout.Cell, headingFunc layout.ListElement) layout.Dimensions {
	// Calculate column widths in pixels. Width is sum of widths.
	w := unit.Value{V: 0, U: colWidths[0].U}
	widths := make([]int, len(colWidths))
	for i, v := range colWidths {
		w.V += v.V
		widths[i] = gtx.Px(v)
	}
	width := gtx.Px(w)
	// Make header position correct scrolling to the far right position.
	if t.AnchorStrategy == Occupy && gtx.Constraints.Max.X < width {
		widths[len(colWidths)-1] += gtx.Px(t.VScrollbarStyle.Width(gtx.Metric))
	}
	// Draw heading.
	headingHeight := drawHeading(gtx, headingFunc, gtx.Px(rowHeightValue), widths, t.State.HorPos)
	// Update constraints and draw grid
	defer op.Offset(f32.Pt(0.0, float32(headingHeight))).Push(gtx.Ops).Pop()
	gtx.Constraints.Max.Y -= headingHeight
	return GridStyle(t).Layout(gtx, rowCount, rowHeightValue, colWidths, cellFunc)
}

// Layout will draw a grid, using fixed column widths and row height.
func (g GridStyle) Layout(gtx layout.Context, rowCount int, rowHeightValue unit.Value,
	colWidths []unit.Value, cellFunc layout.Cell) layout.Dimensions {

	// Determine how much space the scrollbars occupies when present.
	hBarWidth := gtx.Px(g.HScrollbarStyle.Width(gtx.Metric))
	vBarWidth := gtx.Px(g.VScrollbarStyle.Width(gtx.Metric))

	// Calculate column widths in pixels. Width is sum of widths.
	w := unit.Value{V: 0, U: colWidths[0].U}
	widths := make([]int, len(colWidths))
	for i, v := range colWidths {
		w.V += v.V
		widths[i] = gtx.Px(v)
	}

	width := gtx.Px(w)
	// Hide horizontal scroll-bar when not needed.
	if width <= gtx.Constraints.Max.X {
		hBarWidth = 0
		g.State.HorPos = 0
	}

	// Calculate grid size
	rowHeight := gtx.Px(rowHeightValue)
	gridHeight := rowHeight * rowCount

	// Hide vertical scrollbar when not needed
	if gridHeight <= gtx.Constraints.Max.Y-hBarWidth {
		vBarWidth = 0
		g.State.VertPos = 0
	}

	// Reserve space for the scrollbars using the gtx constraints.
	if g.AnchorStrategy == Occupy {
		gtx.Constraints.Max.X -= vBarWidth
		gtx.Constraints.Max.Y -= hBarWidth
	}

	// Make the scroll bar stick to the grid.
	if gtx.Constraints.Max.X > width {
		gtx.Constraints.Max.X = width
	}
	gtx.Constraints.Min = gtx.Constraints.Max

	defer pointer.PassOp{}.Push(gtx.Ops).Pop()

	// Draw grid by running macro clipped to grid size
	cl := clip.Rect{Max: gtx.Constraints.Max}.Push(gtx.Ops)
	dim := g.State.Grid.Layout(gtx, rowCount, rowHeightValue, colWidths, cellFunc)
	cl.Pop()

	// Get horizontal scroll info.
	delta := g.HScrollbarStyle.Scrollbar.ScrollDistance()
	if delta != 0 {
		g.State.HorPos += int(float32(width-vBarWidth) * delta)
		constrain(&g.State.HorPos, 0, width-gtx.Constraints.Max.X)
	}

	// Get vertical scroll info.
	delta = g.VScrollbarStyle.Scrollbar.ScrollDistance()
	if delta != 0 {
		g.State.VertPos += int(math.Round(float64(float32(rowHeight*rowCount) * delta)))
		constrain(&g.State.VertPos, 0, rowHeight*rowCount-gtx.Constraints.Max.Y)
	}

	var Start float32
	var End float32

	// Draw vertical scroll-bar.
	if vBarWidth > 0 {
		c := gtx
		Start = float32(g.State.VertPos) / float32(rowHeight*rowCount)
		End = Start + float32(c.Constraints.Max.Y)/float32(rowHeight*rowCount)
		if g.AnchorStrategy == Overlay {
			c.Constraints.Max.Y -= hBarWidth
		} else {
			c.Constraints.Max.X += vBarWidth
		}
		c.Constraints.Min = c.Constraints.Max
		layout.E.Layout(c, func(gtx layout.Context) layout.Dimensions {
			return g.VScrollbarStyle.Layout(gtx, layout.Vertical, Start, End)
		})
	}

	// Draw horizontal scroll-bar if it is visible.
	if hBarWidth > 0 {
		c := gtx
		Start = float32(g.State.HorPos) / float32(width)
		End = Start + float32(c.Constraints.Max.X)/float32(width)
		if g.AnchorStrategy == Overlay {
			c.Constraints.Max.X -= vBarWidth
		} else {
			c.Constraints.Max.Y += hBarWidth
		}
		c.Constraints.Min = c.Constraints.Max
		layout.S.Layout(c, func(gtx layout.Context) layout.Dimensions {
			return g.HScrollbarStyle.Layout(gtx, layout.Horizontal, Start, End)
		})
	}
	if g.AnchorStrategy == Occupy {
		dim.Size.Y += hBarWidth
	}
	dim.Size.Y += rowHeight
	return dim
}
