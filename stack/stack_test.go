package stack_test

import (
	"testing"

	"github.com/fastygo/base/internal/testutil"
	"github.com/fastygo/base/stack"
)

func TestStackUlTagResolve(t *testing.T) {
	html := testutil.Render(t, stack.Stack(stack.StackProps{Tag: "ul", Class: "items"}))
	testutil.MustContain(t, html, `<ul class="items"`)
	testutil.AssertHeadlessAtoms(t, html)
}
