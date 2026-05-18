package testutil

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func Render(t *testing.T, c interface {
	Render(context.Context, io.Writer) error
}) string {
	t.Helper()
	var buf bytes.Buffer
	if err := c.Render(context.Background(), &buf); err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	return buf.String()
}

// RenderChild wires a single child templ.Component onto ctx before rendering root.
func RenderChild(t *testing.T, root, child templ.Component) string {
	t.Helper()
	ctx := templ.InitializeContext(context.Background())
	ctx = templ.WithChildren(ctx, child)
	var buf bytes.Buffer
	if err := root.Render(ctx, &buf); err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	return buf.String()
}

func MustContain(t *testing.T, html, want string) {
	t.Helper()
	if !strings.Contains(html, want) {
		t.Errorf("expected HTML to contain %q, got:\n%s", want, html)
	}
}

func MustNotContain(t *testing.T, html, notWant string) {
	t.Helper()
	if strings.Contains(html, notWant) {
		t.Errorf("expected HTML NOT to contain %q, got:\n%s", notWant, html)
	}
}

func AssertHeadlessAtoms(t *testing.T, html string) {
	t.Helper()
	MustNotContain(t, html, "ui-")
	MustNotContain(t, html, "bg-")
	MustNotContain(t, html, "text-sm")
	MustNotContain(t, html, "inline-flex")
}
