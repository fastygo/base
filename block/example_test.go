package block_test

import (
	"bytes"
	"context"
	"fmt"

	"github.com/fastygo/base/block"
)

func ExampleBlock() {
	var buf bytes.Buffer
	cmp := block.Block(block.BlockProps{Tag: "section", ID: "hero", Class: "hero-root"})
	if err := cmp.Render(context.Background(), &buf); err != nil {
		panic(err)
	}
	fmt.Print(buf.String())
	// Output: <section class="hero-root" id="hero"></section>
}
