package text_test

import (
	"testing"

	"github.com/fastygo/base/internal/testutil"
	"github.com/fastygo/base/text"
)

func TestTextSpanTagResolve(t *testing.T) {
	html := testutil.Render(t, text.Text(text.TextProps{Tag: "span", Class: "lede"}, "x"))
	testutil.MustContain(t, html, `<span class="lede"`)
	testutil.AssertHeadlessAtoms(t, html)
}
