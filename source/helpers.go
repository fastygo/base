package source

import (
	"strings"

	"github.com/a-h/templ"
)

func sourceAttrs(props SourceProps) templ.Attributes {
	attrs := templ.Attributes{}
	if strings.TrimSpace(props.SrcSet) != "" {
		attrs["srcset"] = props.SrcSet
	}
	if strings.TrimSpace(props.Src) != "" {
		attrs["src"] = props.Src
	}
	if strings.TrimSpace(props.Media) != "" {
		attrs["media"] = props.Media
	}
	if strings.TrimSpace(props.Type) != "" {
		attrs["type"] = props.Type
	}
	if strings.TrimSpace(props.Sizes) != "" {
		attrs["sizes"] = props.Sizes
	}
	return attrs
}
