package block_test

import (
	"testing"

	"github.com/fastygo/base/block"
	"github.com/fastygo/base/internal/testutil"
)

func TestBlockIDAndSemanticTag(t *testing.T) {
	html := testutil.Render(t, block.Block(block.BlockProps{
		ID:    "hero",
		Tag:   "section",
		Class: "landing",
	}))
	testutil.MustContain(t, html, `id="hero"`)
	testutil.MustContain(t, html, `<section`)
	testutil.MustContain(t, html, `class="landing"`)
	testutil.AssertHeadlessAtoms(t, html)
}
