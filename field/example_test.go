package field_test

import (
	"bytes"
	"context"
	"fmt"

	"github.com/fastygo/base/field"
)

func ExampleField_plainInput() {
	var buf bytes.Buffer
	cmp := field.Field(field.FieldProps{Name: "email", Class: "control"})
	if err := cmp.Render(context.Background(), &buf); err != nil {
		panic(err)
	}
	fmt.Print(buf.String())
	// Output: <input id="email" name="email" type="text" class="control" placeholder="" value="" min="" max="" autocomplete="">
}
