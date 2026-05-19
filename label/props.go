package label

import "github.com/a-h/templ"

// LabelProps drives native label wiring.
type LabelProps struct {
	Class   string
	HTMLFor string
	Attrs   templ.Attributes
}
