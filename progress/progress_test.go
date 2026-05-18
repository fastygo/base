package progress_test

import (
	"testing"

	"github.com/fastygo/base/internal/testutil"
	"github.com/fastygo/base/progress"
)

func TestProgressValueMax(t *testing.T) {
	html := testutil.Render(t, progress.Progress(progress.ProgressProps{
		Class: "px",
		Value: "33",
		Max:   "100",
	}))
	testutil.MustContain(t, html, `<progress`)
	testutil.MustContain(t, html, `value="33"`)
	testutil.MustContain(t, html, `max="100"`)
	testutil.AssertHeadlessAtoms(t, html)
}
