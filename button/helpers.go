package button

import (
	"strings"

	"github.com/a-h/templ"
)

func domAttrs(id, role, tabIndex string, attrs templ.Attributes) templ.Attributes {
	out := templ.Attributes{}
	for key, value := range attrs {
		out[key] = value
	}
	if strings.TrimSpace(id) != "" {
		out["id"] = id
	}
	if strings.TrimSpace(role) != "" {
		out["role"] = role
	}
	if strings.TrimSpace(tabIndex) != "" {
		out["tabindex"] = tabIndex
	}
	return out
}

func buttonAttrs(props ButtonProps) templ.Attributes {
	attrs := domAttrs(props.ID, props.Role, props.TabIndex, props.Attrs)
	if strings.TrimSpace(props.AriaLabel) != "" {
		attrs["aria-label"] = props.AriaLabel
	}
	if strings.TrimSpace(props.Href) != "" && props.Disabled {
		attrs["aria-disabled"] = "true"
		attrs["tabindex"] = "-1"
		if _, ok := attrs["role"]; !ok {
			attrs["role"] = "link"
		}
	}
	return attrs
}

func buttonType(t string) string {
	if strings.TrimSpace(t) == "" {
		return "button"
	}
	return t
}
