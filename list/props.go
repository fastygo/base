package list

import "github.com/a-h/templ"

type ListProps struct {
	Class string
	Tag   string
	Attrs templ.Attributes
}

type ListItemProps struct {
	Class string
	Value int
	Attrs templ.Attributes
}
