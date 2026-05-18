package group_test

import (
	"testing"

	"github.com/fastygo/base/group"
	"github.com/fastygo/base/internal/testutil"
)

func TestGroupDataGrow(t *testing.T) {
	html := testutil.Render(t, group.Group(group.GroupProps{Class: "row", Grow: true}))
	testutil.MustContain(t, html, `data-grow`)
	testutil.AssertHeadlessAtoms(t, html)
}
