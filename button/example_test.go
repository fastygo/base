package button_test

import (
	"bytes"
	"context"
	"fmt"

	"github.com/fastygo/base/button"
)

func ExampleButton() {
	var buf bytes.Buffer
	cmp := button.Button(button.ButtonProps{Class: "actions-primary"}, "Save")
	if err := cmp.Render(context.Background(), &buf); err != nil {
		panic(err)
	}
	fmt.Print(buf.String())
	// Output: <button type="button" class="actions-primary">Save</button>
}
