package description_test

import (
	"testing"

	"github.com/fastygo/base/description"
	"github.com/fastygo/base/internal/testutil"
)

func TestDescriptionListStructural(t *testing.T) {
	html := testutil.Render(t, description.DescriptionList(description.DescriptionListProps{Class: "meta"}))
	testutil.MustContain(t, html, `<dl class="meta"`)
	testutil.AssertHeadlessAtoms(t, html)
}
