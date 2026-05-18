package meter_test

import (
	"testing"

	"github.com/fastygo/base/internal/testutil"
	"github.com/fastygo/base/meter"
)

func TestMeterNativeBoundsEmit(t *testing.T) {
	html := testutil.Render(t, meter.Meter(meter.MeterProps{
		Class: "usage",
		Value: "70",
		Min:   "0",
		Max:   "100",
	}))
	testutil.MustContain(t, html, `<meter`)
	testutil.MustContain(t, html, `value="70"`)
	testutil.MustContain(t, html, `min="0"`)
	testutil.MustContain(t, html, `max="100"`)
	testutil.AssertHeadlessAtoms(t, html)
}
