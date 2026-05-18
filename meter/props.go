package meter

import "github.com/a-h/templ"

type MeterProps struct {
	ID      string
	Class   string
	Value   string
	Min     string
	Max     string
	Low     string
	High    string
	Optimum string
	Attrs   templ.Attributes
}
