package output

import "github.com/a-h/templ"

type OutputProps struct {
	ID    string
	Class string
	Name  string
	For   string
	Value string
	Attrs templ.Attributes
}
