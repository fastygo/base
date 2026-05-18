package image_test

import (
	"testing"

	"github.com/fastygo/base/image"
	"github.com/fastygo/base/internal/testutil"
)

func TestImageLoadingDefaultLazy(t *testing.T) {
	html := testutil.Render(t, image.Image(image.ImageProps{
		Src:   "/pic.jpg",
		Alt:   "pic",
		Class: "thumb",
	}))
	testutil.MustContain(t, html, `loading="lazy"`)
	testutil.MustContain(t, html, `src="/pic.jpg"`)
	testutil.AssertHeadlessAtoms(t, html)
}

func TestImageEagerHonored(t *testing.T) {
	html := testutil.Render(t, image.Image(image.ImageProps{Src: "/p", Alt: "a", Loading: "eager"}))
	testutil.MustContain(t, html, `loading="eager"`)
}
