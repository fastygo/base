package form_test

import (
	"testing"

	"github.com/a-h/templ"

	"github.com/fastygo/base/form"
	"github.com/fastygo/base/internal/testutil"
)

func TestFormNativeAttrsSlot(t *testing.T) {
	html := testutil.Render(t, form.Form(form.FormProps{
		Class:  "signup",
		Method: "post",
		Action: "/signup",
		Attrs:  templ.Attributes{"data-test": "1"},
	}))
	testutil.MustContain(t, html, `method="post"`)
	testutil.MustContain(t, html, `action="/signup"`)
	testutil.MustContain(t, html, `data-test="1"`)
	testutil.AssertHeadlessAtoms(t, html)
}
