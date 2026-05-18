package datalist_test

import (
	"testing"

	"github.com/fastygo/base/datalist"
	"github.com/fastygo/base/internal/testutil"
)

func TestDataListLinkageID(t *testing.T) {
	html := testutil.Render(t, datalist.DataList(datalist.DataListProps{
		ID:    "opts",
		Class: "hints",
	}))
	testutil.MustContain(t, html, `id="opts"`)
	testutil.AssertHeadlessAtoms(t, html)
}
