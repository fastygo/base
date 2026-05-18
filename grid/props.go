package grid

type GridProps struct {
	Class string
	Cols  string
}

type GridColProps struct {
	Class string
	Span  int
	Start int
	End   int
	Order int
}
