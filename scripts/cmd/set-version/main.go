package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: go run ./scripts/cmd/set-version <version>")
		os.Exit(1)
	}

	version := os.Args[1]
	src, err := os.ReadFile("doc.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "read doc.go: %v\n", err)
		os.Exit(1)
	}

	re := regexp.MustCompile(`const Version = "\d+\.\d+\.\d+"`)
	replacement := fmt.Sprintf(`const Version = "%s"`, version)
	updated := re.ReplaceAllString(string(src), replacement)
	if updated == string(src) {
		fmt.Fprintln(os.Stderr, `version constant not found in doc.go (expected line: const Version = "0.1.0")`)
		os.Exit(1)
	}

	if err := os.WriteFile("doc.go", []byte(updated), 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "write doc.go: %v\n", err)
		os.Exit(1)
	}
}
