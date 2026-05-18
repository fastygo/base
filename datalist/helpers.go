package datalist

import (
	"strings"

	"github.com/a-h/templ"
)

func dataListAttrs(p DataListProps) templ.Attributes {
	attrs := templ.Attributes{}
	if strings.TrimSpace(p.ID) != "" {
		attrs["id"] = p.ID
	}
	return attrs
}
