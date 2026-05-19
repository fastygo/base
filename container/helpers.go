package container

import (
	"github.com/a-h/templ"
	"github.com/fastygo/base/utils"
)

func containerSpreadAttrs(props ContainerProps) templ.Attributes {
	return utils.MergeAttrs(templ.Attributes{}, props.Attrs)
}
