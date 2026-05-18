package fieldset_test

import (
	"testing"

	"github.com/fastygo/base/fieldset"
	"github.com/fastygo/base/internal/testutil"
)

func TestFieldsetDisabled(t *testing.T) {
	html := testutil.Render(t, fieldset.Fieldset(fieldset.FieldsetProps{
		Class:    "fs",
		Disabled: true,
	}))
	testutil.MustContain(t, html, `<fieldset`)
	testutil.MustContain(t, html, `disabled`)
	testutil.AssertHeadlessAtoms(t, html)
}
