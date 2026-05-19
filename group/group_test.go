package group_test

import (
	"testing"

	"github.com/a-h/templ"
	"github.com/fastygo/base/group"
	"github.com/fastygo/base/internal/testutil"
)

func TestGroupSpreadAttrs(t *testing.T) {
	html := testutil.Render(t, group.Group(group.GroupProps{
		Class: "row w-full",
		Attrs: templ.Attributes{"data-test": "g"},
	}))
	testutil.MustContain(t, html, `data-test="g"`)
	testutil.MustContain(t, html, `class="row w-full"`)
	testutil.AssertHeadlessAtoms(t, html)
}
