package legend_test

import (
	"testing"

	"github.com/fastygo/base/internal/testutil"
	"github.com/fastygo/base/legend"
)

func TestLegendStructural(t *testing.T) {
	html := testutil.Render(t, legend.Legend(legend.LegendProps{Class: "cap"}))
	testutil.MustContain(t, html, `<legend class="cap"`)
	testutil.AssertHeadlessAtoms(t, html)
}
