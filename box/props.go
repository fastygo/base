package box

import "github.com/a-h/templ"

// BoxProps controls tag selection and styling slot Class.
type BoxProps struct {
	Class string
	Tag   string
	Attrs templ.Attributes
}
