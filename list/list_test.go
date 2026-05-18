package list_test

import (
	"testing"

	"github.com/fastygo/base/internal/testutil"
	"github.com/fastygo/base/list"
)

func TestListOrderedTagResolve(t *testing.T) {
	html := testutil.Render(t, list.List(list.ListProps{Tag: "ol", Class: "steps"}))
	testutil.MustContain(t, html, `<ol class="steps"`)
	testutil.AssertHeadlessAtoms(t, html)
}

func TestListItemValueAttr(t *testing.T) {
	html := testutil.Render(t, list.ListItem(list.ListItemProps{Class: "it", Value: 3}))
	testutil.MustContain(t, html, `value="3"`)
}
