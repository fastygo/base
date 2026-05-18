package icon_test

import (
	"testing"

	"github.com/fastygo/base/icon"
	"github.com/fastygo/base/internal/testutil"
)

func TestIconDataNameSlot(t *testing.T) {
	html := testutil.Render(t, icon.Icon(icon.IconProps{Name: "close", Class: "glyph"}))
	testutil.MustContain(t, html, `data-icon="close"`)
	testutil.MustContain(t, html, `class="glyph"`)
	testutil.AssertHeadlessAtoms(t, html)
}
