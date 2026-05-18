package block

import (
	"strings"

	"github.com/a-h/templ"
)

func blockAttrs(id string) templ.Attributes {
	attrs := templ.Attributes{}
	if strings.TrimSpace(id) != "" {
		attrs["id"] = id
	}
	return attrs
}
