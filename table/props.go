package table

import "github.com/a-h/templ"

type TableProps struct {
	Class string
	Attrs templ.Attributes
}

type TableCaptionProps struct {
	Class string
}

type TableSectionProps struct {
	Class string
}

type TableRowProps struct {
	Class string
}

type TableCellProps struct {
	Class   string
	Scope   string
	ColSpan int
	RowSpan int
	Headers string
	Abbr    string
}

type TableColGroupProps struct {
	Class string
	Span  int
}

type TableColProps struct {
	Class string
	Span  int
}
