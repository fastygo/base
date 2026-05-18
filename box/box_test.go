package box_test

import (
	"testing"

	"github.com/fastygo/base/box"
	"github.com/fastygo/base/internal/testutil"
)

func TestBoxDefaultDivWithClassSlot(t *testing.T) {
	html := testutil.Render(t, box.Box(box.BoxProps{Class: "x-root"}))
	testutil.MustContain(t, html, `<div class="x-root"`)
	testutil.AssertHeadlessAtoms(t, html)
}

func TestBoxSectionTagResolve(t *testing.T) {
	html := testutil.Render(t, box.Box(box.BoxProps{Tag: "section"}))
	testutil.MustContain(t, html, `<section`)
}

func TestBoxInvalidTagFallsBackToDiv(t *testing.T) {
	html := testutil.Render(t, box.Box(box.BoxProps{Tag: "canvas"}))
	testutil.MustContain(t, html, `<div`)
}
