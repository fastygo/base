package table_test

import (
	"bytes"
	"context"
	"fmt"

	"github.com/fastygo/base/table"
)

func ExampleTable_plain() {
	var buf bytes.Buffer
	cmp := table.Table(table.TableProps{Class: "data-grid"})
	if err := cmp.Render(context.Background(), &buf); err != nil {
		panic(err)
	}
	fmt.Print(buf.String())
	// Output: <table class="data-grid"></table>
}
