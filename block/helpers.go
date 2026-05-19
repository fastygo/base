package block

import (
	"strings"

	"github.com/a-h/templ"

	"github.com/fastygo/base/utils"
)

func blockAttrs(id string) templ.Attributes {
	attrs := templ.Attributes{}
	if strings.TrimSpace(id) != "" {
		attrs["id"] = id
	}
	return attrs
}

func blockSpreadAttrs(p BlockProps) templ.Attributes {
	return utils.MergeAttrs(blockAttrs(p.ID), p.Attrs)
}
