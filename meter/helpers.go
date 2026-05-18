package meter

import (
	"strings"

	"github.com/a-h/templ"
)

func meterAttrs(p MeterProps) templ.Attributes {
	attrs := templ.Attributes{}
	for key, value := range p.Attrs {
		attrs[key] = value
	}
	if strings.TrimSpace(p.ID) != "" {
		attrs["id"] = p.ID
	}
	if strings.TrimSpace(p.Value) != "" {
		attrs["value"] = p.Value
	}
	if strings.TrimSpace(p.Min) != "" {
		attrs["min"] = p.Min
	}
	if strings.TrimSpace(p.Max) != "" {
		attrs["max"] = p.Max
	}
	if strings.TrimSpace(p.Low) != "" {
		attrs["low"] = p.Low
	}
	if strings.TrimSpace(p.High) != "" {
		attrs["high"] = p.High
	}
	if strings.TrimSpace(p.Optimum) != "" {
		attrs["optimum"] = p.Optimum
	}
	return attrs
}
