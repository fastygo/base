package block

import "github.com/a-h/templ"

type BlockProps struct {
	ID    string
	Class string
	Tag   string
	Attrs templ.Attributes
}
