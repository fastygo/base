package stack

import (
	"github.com/a-h/templ"
	"github.com/fastygo/base/utils"
)

func stackSpreadAttrs(props StackProps) templ.Attributes {
	return utils.MergeAttrs(templ.Attributes{}, props.Attrs)
}
