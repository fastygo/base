package optgroup

import (
	"strings"

	"github.com/a-h/templ"
)

func optGroupAttrs(p OptGroupProps) templ.Attributes {
	attrs := templ.Attributes{}
	if strings.TrimSpace(p.Label) != "" {
		attrs["label"] = p.Label
	}
	if p.Disabled {
		attrs["disabled"] = true
	}
	return attrs
}
