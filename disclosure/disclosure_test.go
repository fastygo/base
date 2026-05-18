package disclosure_test

import (
	"testing"

	"github.com/fastygo/base/disclosure"
	"github.com/fastygo/base/internal/testutil"
)

func TestDisclosureNativeOpenToggle(t *testing.T) {
	html := testutil.Render(t, disclosure.Disclosure(disclosure.DisclosureProps{Open: true, Class: "acc"}))
	testutil.MustContain(t, html, `<details`)
	testutil.MustContain(t, html, `open`)
	testutil.AssertHeadlessAtoms(t, html)
}
