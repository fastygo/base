package figure_test

import (
	"testing"

	"github.com/fastygo/base/figure"
	"github.com/fastygo/base/internal/testutil"
)

func TestFigureCaptionSlots(t *testing.T) {
	html := testutil.Render(t, figure.FigureCaption(figure.FigureCaptionProps{Class: "cap"}))
	testutil.MustContain(t, html, `<figcaption class="cap"`)
	testutil.AssertHeadlessAtoms(t, html)
}
