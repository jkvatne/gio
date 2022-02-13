// SPDX-License-Identifier: Unlicense OR MIT

package layout

import (
	"image"

	"gioui.org/f32"
	"gioui.org/gesture"
	"gioui.org/io/pointer"
	"gioui.org/op"
	"gioui.org/op/clip"
)

type Grid struct {
	HorPos  int
	VertPos int
	Vscroll gesture.Scroll
	Hscroll gesture.Scroll
	dims    []Dimensions
	call    []op.CallOp
}

// Cell is the layout function for a grid cell, with row,col parameters.
type Cell func(gtx Context, col int, row int) Dimensions

// FindColStart find first and last visible cells in row by looping over all cells.
// Also returns offset into first cell (in pixels).
func FindColStart(HorPos int, maxX int, widths []int) (firstCol int, lastCol int, colOffset int) {
	sum := 0
	for _, w := range widths {
		sum += w
	}
	if HorPos > sum-maxX {
		HorPos = sum - maxX
	}
	if HorPos < 0 {
		HorPos = 0
	}
	pos := 0
	for i, w := range widths {
		if pos <= HorPos {
			firstCol = i
			colOffset = pos
		} else if pos > HorPos+maxX {
			lastCol = i - 1
			break
		}
		lastCol = i
		pos += w
	}
	colOffset -= HorPos
	return
}

// FindRowStart find first and last visible row in row by looping over all cells.
// Also returns offset which is the offset into first row.
func FindRowStart(VertPos int, maxY int, rowCount int, rowHeight int) (firstRow int, lastRow int, rowOffset int) {
	visibleRows := maxY/rowHeight + 1
	firstRow = VertPos / rowHeight
	lastRow = firstRow + visibleRows
	rowOffset = firstRow*rowHeight - VertPos
	if lastRow > rowCount-1 {
		lastRow = rowCount - 1
		if VertPos > 0 {
			firstRow = rowCount - visibleRows
		}
		if firstRow > 0 {
			rowOffset = maxY - visibleRows*rowHeight
		}
	}
	if visibleRows > rowCount {
		rowOffset = 0
	}
	if firstRow < 0 {
		firstRow = 0
	}
	return
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

// Layout the Grid.
func (g *Grid) Layout(gtx Context, rowCount int, rowHeight int, widths []int, cellFunc Cell) Dimensions {
	listDims := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)

	width := 0
	for _, w := range widths {
		width += w
	}

	// Update horizontal scroll position.
	hScrollDelta := g.Hscroll.Scroll(gtx.Metric, gtx, gtx.Now, gesture.Horizontal)
	if hScrollDelta != 0 {
		g.HorPos += hScrollDelta
		constrain(&g.HorPos, 0, width-gtx.Constraints.Max.X)
	}

	// Get vertical scroll info.
	vScrollDelta := g.Vscroll.Scroll(gtx.Metric, gtx, gtx.Now, gesture.Vertical)
	if vScrollDelta != 0 {
		g.VertPos += vScrollDelta
		constrain(&g.VertPos, 0, rowHeight*rowCount-gtx.Constraints.Max.Y)
	}

	firstRow, lastRow, rowOffset := FindRowStart(g.VertPos, gtx.Constraints.Max.Y, rowCount, rowHeight)
	firstCol, lastCol, colOffset := FindColStart(g.HorPos, gtx.Constraints.Max.X, widths)
	// Draw rows into macro.
	macro := op.Record(gtx.Ops)
	clp := clip.Rect{Max: gtx.Constraints.Max}.Push(gtx.Ops)
	for row := firstRow; row <= lastRow; row++ {
		// Lay out visible cells in a row.
		g.dims = g.dims[:0]
		g.call = g.call[:0]
		for col := firstCol; col <= lastCol; col++ {
			c := gtx
			c.Constraints = Exact(image.Pt(widths[col], rowHeight))
			// Draw the cell into macro
			macro := op.Record(c.Ops)
			g.dims = append(g.dims, cellFunc(c, col, row))
			g.call = append(g.call, macro.Stop())
		}
		// Generate all the rendering commands for the children, translated to correct location.
		xPos := colOffset
		for col := firstCol; col <= lastCol; col++ {
			trans := op.Offset(f32.Pt(float32(xPos), float32(rowOffset))).Push(gtx.Ops)
			clp := clip.Rect{Max: g.dims[col-firstCol].Size}.Push(gtx.Ops)
			g.call[col-firstCol].Add(gtx.Ops)
			clp.Pop()
			trans.Pop()
			xPos += g.dims[col-firstCol].Size.X
		}
		rowOffset += rowHeight
	}
	clp.Pop()
	call := macro.Stop()

	// Draw grid by running macro clipped to grid size
	defer pointer.PassOp{}.Push(gtx.Ops).Pop()
	c := gtx.Constraints.Max
	c.Y = gtx.Constraints.Max.Y
	c.X = gtx.Constraints.Max.X
	cl := clip.Rect{Max: c}.Push(gtx.Ops)
	call.Add(gtx.Ops)
	cl.Pop()

	// Enable scroll wheel within the grid
	c = gtx.Constraints.Max
	c.Y = gtx.Constraints.Max.Y
	c.X = gtx.Constraints.Max.X
	cl = clip.Rect{Max: c}.Push(gtx.Ops)
	g.Vscroll.Add(gtx.Ops, image.Rect(0, -c.Y/2, 0, c.Y/2))
	g.Hscroll.Add(gtx.Ops, image.Rect(-c.X, 0, c.X, 0))
	cl.Pop()

	return Dimensions{Size: listDims, Baseline: 0}
}
