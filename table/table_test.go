package table_test

import (
	"testing"

	"github.com/fastygo/base/internal/testutil"
	"github.com/fastygo/base/table"
)

func TestTableHeadScopeCol(t *testing.T) {
	html := testutil.Render(t, table.TableHeadCell(table.TableCellProps{
		Class: "hc",
		Scope: "col",
	}))
	testutil.MustContain(t, html, `scope="col"`)
	testutil.AssertHeadlessAtoms(t, html)
}

func TestTableCellColSpan(t *testing.T) {
	html := testutil.Render(t, table.TableCell(table.TableCellProps{
		Class:   "c",
		ColSpan: 2,
		RowSpan: 3,
		Headers: "h1",
	}))
	testutil.MustContain(t, html, `colspan="2"`)
	testutil.MustContain(t, html, `rowspan="3"`)
	testutil.MustContain(t, html, `headers="h1"`)
	testutil.AssertHeadlessAtoms(t, html)
}

func TestColGroupSpan(t *testing.T) {
	html := testutil.Render(t, table.TableColGroup(table.TableColGroupProps{
		Class: "cg",
		Span:  2,
	}))
	testutil.MustContain(t, html, `span="2"`)
	testutil.AssertHeadlessAtoms(t, html)
}
