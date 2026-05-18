package field

import (
	"strconv"
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

func fieldAttrs(p FieldProps) templ.Attributes {
	attrs := domAttrs(p.ID, p.Role, p.TabIndex, p.Attrs)
	delete(attrs, "id")
	if strings.TrimSpace(p.AriaLabel) != "" {
		attrs["aria-label"] = p.AriaLabel
	}
	if p.Type == "checkbox" && p.Switch {
		attrs["role"] = "switch"
		attrs["aria-checked"] = strconv.FormatBool(p.Checked)
	}
	return attrs
}

func fieldRows(rows int) int {
	if rows <= 0 {
		return 4
	}
	return rows
}

func fieldInputType(t string) string {
	if strings.TrimSpace(t) == "" {
		return "text"
	}
	return t
}

func fieldID(p FieldProps) string {
	if strings.TrimSpace(p.ID) != "" {
		return p.ID
	}
	if strings.TrimSpace(p.Name) != "" {
		return p.Name
	}
	return ""
}
