package list

import (
	"strconv"

	"github.com/a-h/templ"
	"github.com/fastygo/base/utils"
)

func listTag(tag string) string {
	return utils.ResolveTag(tag, "ul", utils.TagGroupList)
}

func listItemAttrs(p ListItemProps) templ.Attributes {
	if p.Value <= 0 {
		return templ.Attributes{}
	}
	return templ.Attributes{"value": strconv.Itoa(p.Value)}
}
