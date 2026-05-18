package optgroup_test

import (
	"testing"

	"github.com/fastygo/base/internal/testutil"
	"github.com/fastygo/base/optgroup"
)

func TestOptGroupLabelAttribute(t *testing.T) {
	html := testutil.Render(t, optgroup.OptGroup(optgroup.OptGroupProps{
		Label: "Top",
		Class: "grp",
	}))
	testutil.MustContain(t, html, `label="Top"`)
	testutil.AssertHeadlessAtoms(t, html)
}
