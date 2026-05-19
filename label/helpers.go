package label

import (
	"strings"

	"github.com/a-h/templ"
	"github.com/fastygo/base/utils"
)

func labelSpreadAttrs(props LabelProps) templ.Attributes {
	out := utils.MergeAttrs(templ.Attributes{}, props.Attrs)
	if strings.TrimSpace(props.HTMLFor) != "" {
		out = utils.MergeAttrs(out, templ.Attributes{"for": props.HTMLFor})
	}
	return out
}
