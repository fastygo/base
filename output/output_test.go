package output_test

import (
	"testing"

	"github.com/fastygo/base/internal/testutil"
	"github.com/fastygo/base/output"
)

func TestOutputForLinkage(t *testing.T) {
	html := testutil.Render(t, output.Output(output.OutputProps{
		Name:  "sum",
		Value: "42",
		For:   "a,b",
		Class: "ro",
	}))
	testutil.MustContain(t, html, `for="a,b"`)
	testutil.MustContain(t, html, `42`)
	testutil.AssertHeadlessAtoms(t, html)
}
