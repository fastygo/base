package field_test

import (
	"testing"

	"github.com/fastygo/base/field"
	"github.com/fastygo/base/internal/testutil"
)

func TestFieldInputPlainClassSlot(t *testing.T) {
	html := testutil.Render(t, field.Field(field.FieldProps{
		Name:  "email",
		Class: "control",
	}))
	testutil.MustContain(t, html, `name="email"`)
	testutil.MustContain(t, html, `class="control"`)
	testutil.MustContain(t, html, `type="text"`)
	testutil.AssertHeadlessAtoms(t, html)
}

func TestFieldSwitchSetsRole(t *testing.T) {
	html := testutil.Render(t, field.Field(field.FieldProps{
		Name:    "tone",
		Type:    "checkbox",
		Class:   "switch",
		Switch:  true,
		Checked: true,
	}))
	testutil.MustContain(t, html, `role="switch"`)
	testutil.MustContain(t, html, `aria-checked="true"`)
}

func TestFieldLabeledUsesWrapperClasses(t *testing.T) {
	html := testutil.Render(t, field.Field(field.FieldProps{
		Name:         "bio",
		Label:        "Bio",
		WrapperClass: "wrap",
		LabelClass:   "lbl",
		Class:        "inp",
	}))
	testutil.MustContain(t, html, `class="wrap"`)
	testutil.MustContain(t, html, `<label`)
	testutil.MustContain(t, html, `class="lbl"`)
	testutil.MustContain(t, html, `<input`)
	testutil.AssertHeadlessAtoms(t, html)
}
