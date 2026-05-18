package source_test

import (
	"testing"

	"github.com/fastygo/base/internal/testutil"
	"github.com/fastygo/base/source"
)

func TestSourceSrcSet(t *testing.T) {
	html := testutil.Render(t, source.Source(source.SourceProps{
		SrcSet: "/small.jpg 640w",
		Type:   "image/jpeg",
	}))
	testutil.MustContain(t, html, `srcset="/small.jpg 640w"`)
	testutil.MustContain(t, html, `type="image/jpeg"`)
	testutil.AssertHeadlessAtoms(t, html)
}
