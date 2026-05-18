package integration_test

import (
	"testing"

	"github.com/fastygo/base/block"
	"github.com/fastygo/base/button"
	"github.com/fastygo/base/field"
	"github.com/fastygo/base/internal/testutil"
	"github.com/fastygo/base/stack"
)

func TestComposeBlockWrapsStackAcrossPackages(t *testing.T) {
	html := testutil.RenderChild(t,
		block.Block(block.BlockProps{Tag: "section", Class: "page"}),
		stack.Stack(stack.StackProps{Tag: "div", Class: "col"}),
	)
	testutil.MustContain(t, html, `<section class="page"`)
	testutil.MustContain(t, html, `<div class="col"`)
	testutil.AssertHeadlessAtoms(t, html)
}

func TestFieldAndButtonCoexistImports(t *testing.T) {
	html := testutil.Render(t, field.Field(field.FieldProps{Name: "q"}))
	testutil.MustContain(t, html, `<input`)

	html2 := testutil.Render(t, button.Button(button.ButtonProps{}, "Go"))
	testutil.MustContain(t, html2, `<button`)
	testutil.AssertHeadlessAtoms(t, html2)
}
