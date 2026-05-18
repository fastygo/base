package fieldset

import "github.com/a-h/templ"

type FieldsetProps struct {
	Class    string
	Name     string
	Form     string
	Disabled bool
	Attrs    templ.Attributes
}
