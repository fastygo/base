package grid_test

import (
	"testing"

	"github.com/fastygo/base/grid"
	"github.com/fastygo/base/internal/testutil"
)

func TestGridDataColsAnchor(t *testing.T) {
	html := testutil.Render(t, grid.Grid(grid.GridProps{Class: "g", Cols: "1-4"}))
	testutil.MustContain(t, html, `data-cols="1-4"`)
	testutil.AssertHeadlessAtoms(t, html)
}

func TestGridColDataSpan(t *testing.T) {
	html := testutil.Render(t, grid.GridCol(grid.GridColProps{Class: "c", Span: 2, Start: 1}))
	testutil.MustContain(t, html, `data-span="2"`)
	testutil.MustContain(t, html, `data-start="1"`)
	testutil.AssertHeadlessAtoms(t, html)
}
