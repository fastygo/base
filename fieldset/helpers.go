package fieldset

import (
	"strings"

	"github.com/a-h/templ"
)

func fieldsetAttrs(p FieldsetProps) templ.Attributes {
	attrs := templ.Attributes{}
	for key, value := range p.Attrs {
		attrs[key] = value
	}
	if strings.TrimSpace(p.Name) != "" {
		attrs["name"] = p.Name
	}
	if strings.TrimSpace(p.Form) != "" {
		attrs["form"] = p.Form
	}
	if p.Disabled {
		attrs["disabled"] = true
	}
	return attrs
}
