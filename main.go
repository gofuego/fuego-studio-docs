package main

import (
	"fmt"
	"os"

	doctheme "github.com/gofuego/fuego-doctheme"
	"github.com/gofuego/fuego/engine"
	"github.com/gofuego/fuego/parsers/markdown"
)

// fuego-studio-docs is the public, user-facing documentation for fuego-studio:
// how to host and edit GitHub-backed sites. It is a Fuego site using the shared
// public docs theme; tutorials, how-to guides, reference, and explanation follow
// the Diátaxis structure.
func main() {
	eng := engine.New()
	eng.Register(markdown.Parser())
	eng.Use(doctheme.Public())

	if err := eng.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "fuego: %v\n", err)
		os.Exit(1)
	}
}
