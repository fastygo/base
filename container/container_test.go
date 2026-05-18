package container_test

import (
	"testing"

	"github.com/fastygo/base/container"
	"github.com/fastygo/base/internal/testutil"
)

func TestContainerMainSlot(t *testing.T) {
	html := testutil.Render(t, container.Container(container.ContainerProps{Tag: "main", Class: "shell"}))
	testutil.MustContain(t, html, `<main class="shell"`)
	testutil.AssertHeadlessAtoms(t, html)
}
