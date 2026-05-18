package output

import (
	"strings"

	"github.com/a-h/templ"
)

func outputAttrs(p OutputProps) templ.Attributes {
	attrs := templ.Attributes{}
	for key, value := range p.Attrs {
		attrs[key] = value
	}
	if strings.TrimSpace(p.ID) != "" {
		attrs["id"] = p.ID
	}
	if strings.TrimSpace(p.Name) != "" {
		attrs["name"] = p.Name
	}
	if strings.TrimSpace(p.For) != "" {
		attrs["for"] = p.For
	}
	return attrs
}
