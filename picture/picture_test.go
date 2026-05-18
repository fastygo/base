package picture_test

import (
	"testing"

	"github.com/fastygo/base/internal/testutil"
	"github.com/fastygo/base/picture"
)

func TestPictureStructural(t *testing.T) {
	html := testutil.Render(t, picture.Picture(picture.PictureProps{Class: "ph"}))
	testutil.MustContain(t, html, `<picture class="ph"`)
	testutil.AssertHeadlessAtoms(t, html)
}
