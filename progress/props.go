package progress

import "github.com/a-h/templ"

type ProgressProps struct {
	ID    string
	Class string
	Value string
	Max   string
	Attrs templ.Attributes
}
