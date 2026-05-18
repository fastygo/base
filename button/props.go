package button

import "github.com/a-h/templ"

// ButtonProps maps to native behaviors; styling is Caller-controlled via Class and Attrs.
type ButtonProps struct {
	ID        string
	Role      string
	TabIndex  string
	Attrs     templ.Attributes
	Href      string
	Class     string
	Type      string
	AriaLabel string
	Disabled  bool
}
