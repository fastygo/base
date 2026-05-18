package form

import (
	"strings"

	"github.com/a-h/templ"
)

func domAttrsID(id string, attrs templ.Attributes) templ.Attributes {
	out := templ.Attributes{}
	for key, value := range attrs {
		out[key] = value
	}
	if strings.TrimSpace(id) != "" {
		out["id"] = id
	}
	return out
}

func formAttrs(p FormProps) templ.Attributes {
	attrs := domAttrsID(p.ID, p.Attrs)
	if strings.TrimSpace(p.Action) != "" {
		attrs["action"] = p.Action
	}
	if strings.TrimSpace(p.Method) != "" {
		attrs["method"] = p.Method
	}
	if strings.TrimSpace(p.Enctype) != "" {
		attrs["enctype"] = p.Enctype
	}
	if strings.TrimSpace(p.Autocomplete) != "" {
		attrs["autocomplete"] = p.Autocomplete
	}
	if strings.TrimSpace(p.Name) != "" {
		attrs["name"] = p.Name
	}
	if strings.TrimSpace(p.Target) != "" {
		attrs["target"] = p.Target
	}
	if p.NoValidate {
		attrs["novalidate"] = true
	}
	return attrs
}
