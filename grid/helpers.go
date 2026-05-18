package grid

import (
	"strconv"
	"strings"

	"github.com/a-h/templ"
)

func gridDataAttrs(cols string) templ.Attributes {
	if strings.TrimSpace(cols) == "" {
		return templ.Attributes{}
	}
	return templ.Attributes{"data-cols": cols}
}

func gridColDataAttrs(p GridColProps) templ.Attributes {
	a := templ.Attributes{}
	if p.Span > 0 {
		a["data-span"] = strconv.Itoa(p.Span)
	}
	if p.Start > 0 {
		a["data-start"] = strconv.Itoa(p.Start)
	}
	if p.End > 0 {
		a["data-end"] = strconv.Itoa(p.End)
	}
	if p.Order > 0 {
		a["data-order"] = strconv.Itoa(p.Order)
	}
	return a
}
