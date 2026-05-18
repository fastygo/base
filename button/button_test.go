package button_test

import (
	"testing"

	"github.com/fastygo/base/button"
	"github.com/fastygo/base/internal/testutil"
)

func TestButtonNative(t *testing.T) {
	html := testutil.Render(t, button.Button(button.ButtonProps{}, "Go"))
	testutil.MustContain(t, html, `<button`)
	testutil.MustContain(t, html, `type="button"`)
	testutil.MustContain(t, html, `Go`)
	testutil.AssertHeadlessAtoms(t, html)
}

func TestButtonLinkHref(t *testing.T) {
	html := testutil.Render(t, button.Button(button.ButtonProps{Href: "/dash"}, "Link"))
	testutil.MustContain(t, html, `<a`)
	testutil.MustContain(t, html, `href="/dash"`)
	testutil.AssertHeadlessAtoms(t, html)
}

func TestButtonDisabledLinkUsesAriaDisabled(t *testing.T) {
	html := testutil.Render(t, button.Button(button.ButtonProps{
		Href:     "/blocked",
		Disabled: true,
	}, "Nope"))
	testutil.MustContain(t, html, `aria-disabled="true"`)
	testutil.MustContain(t, html, `tabindex="-1"`)
	testutil.MustContain(t, html, `rel="nofollow noopener noreferrer"`)
}
