package title_test

import (
	"testing"

	"github.com/fastygo/base/internal/testutil"
	"github.com/fastygo/base/title"
)

func TestTitleHeadingOrderSlot(t *testing.T) {
	html := testutil.Render(t, title.Title(title.TitleProps{Order: 1, Class: "hd"}, "Hi"))
	testutil.MustContain(t, html, `<h1 class="hd"`)
	testutil.MustContain(t, html, `Hi`)
	testutil.AssertHeadlessAtoms(t, html)
}
