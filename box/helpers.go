package box

import (
	"github.com/a-h/templ"
	"github.com/fastygo/base/utils"
)

func boxSpreadAttrs(props BoxProps) templ.Attributes {
	return utils.MergeAttrs(templ.Attributes{}, props.Attrs)
}
