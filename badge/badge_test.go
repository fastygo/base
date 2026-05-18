package badge_test

import (
	"testing"

	"github.com/fastygo/base/badge"
	"github.com/fastygo/base/internal/testutil"
)

func TestBadgeSpansClassSlot(t *testing.T) {
	html := testutil.Render(t, badge.Badge(badge.BadgeProps{Class: "chip"}, "v1"))
	testutil.MustContain(t, html, `<span class="chip"`)
	testutil.MustContain(t, html, `v1`)
	testutil.AssertHeadlessAtoms(t, html)
}
