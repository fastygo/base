package field

import "github.com/a-h/templ"

// FieldOption is a select pairing.
type FieldOption struct {
	Value string
	Label string
}

// FieldProps drives native controls; Caller supplies WrapperClass/LabelClass/HintClass/ErrorClass presets.
type FieldProps struct {
	ID           string
	Role         string
	TabIndex     string
	Attrs        templ.Attributes
	Class        string
	Type         string
	Name         string
	Placeholder  string
	Value        string
	Rows         int
	Min          string
	Max          string
	Checked      bool
	Disabled     bool
	Required     bool
	Autocomplete string
	Component    string
	Options      []FieldOption
	Label        string
	Hint         string
	Error        string
	AriaLabel    string
	Switch       bool
	WrapperClass string
	LabelClass   string
	HintClass    string
	ErrorClass   string
}
