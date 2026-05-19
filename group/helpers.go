package group

import (
	"github.com/a-h/templ"
	"github.com/fastygo/base/utils"
)

func groupSpreadAttrs(props GroupProps) templ.Attributes {
	return utils.MergeAttrs(templ.Attributes{}, props.Attrs)
}
