package table

import (
	"strconv"
	"strings"

	"github.com/a-h/templ"
)

func mergeBaseAttrs(extra templ.Attributes) templ.Attributes {
	attrs := templ.Attributes{}
	for k, v := range extra {
		attrs[k] = v
	}
	return attrs
}

func tableAttrs(p TableProps) templ.Attributes {
	return mergeBaseAttrs(p.Attrs)
}

func tableCellAttrs(p TableCellProps, heading bool) templ.Attributes {
	attrs := templ.Attributes{}
	if heading {
		scope := tableScope(p.Scope)
		if scope != "" {
			attrs["scope"] = scope
		}
		if strings.TrimSpace(p.Abbr) != "" {
			attrs["abbr"] = p.Abbr
		}
	}
	if p.ColSpan > 0 {
		attrs["colspan"] = strconv.Itoa(p.ColSpan)
	}
	if p.RowSpan > 0 {
		attrs["rowspan"] = strconv.Itoa(p.RowSpan)
	}
	if strings.TrimSpace(p.Headers) != "" {
		attrs["headers"] = p.Headers
	}
	return attrs
}

func tableScope(value string) string {
	switch strings.TrimSpace(value) {
	case "row", "col", "rowgroup", "colgroup":
		return strings.TrimSpace(value)
	default:
		return ""
	}
}

func spanAttrs(span int) templ.Attributes {
	if span <= 0 {
		return templ.Attributes{}
	}
	return templ.Attributes{"span": strconv.Itoa(span)}
}
