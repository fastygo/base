package form

import "github.com/a-h/templ"

// FormProps covers native HTML form semantics.
type FormProps struct {
	ID           string
	Class        string
	Action       string
	Method       string
	Enctype      string
	Autocomplete string
	Name         string
	Target       string
	NoValidate   bool
	Attrs        templ.Attributes
}
