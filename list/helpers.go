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

func listItemMergedAttrs(p ListItemProps) templ.Attributes {
	return utils.MergeAttrs(listItemAttrs(p), p.Attrs)
}

func listSpreadAttrs(props ListProps) templ.Attributes {
	return utils.MergeAttrs(templ.Attributes{}, props.Attrs)
}
