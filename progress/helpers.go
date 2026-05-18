package progress

import (
	"strings"

	"github.com/a-h/templ"
)

func progressAttrs(p ProgressProps) templ.Attributes {
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
	if strings.TrimSpace(p.Max) != "" {
		attrs["max"] = p.Max
	}
	return attrs
}
