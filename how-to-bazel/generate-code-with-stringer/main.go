package main

import (
	"fmt"

	"github.com/ekhabarov/blog-code-snippets/how-to-bazel/generate-code-with-stringer/cst"
)

func main() {
	fmt.Printf("constant: %s\n", cst.ConstA)
	fmt.Printf("constant: %s\n", cst.ConstB)
	fmt.Printf("constant: %s\n", cst.ConstC)
	// Output :
	// constant: ConstA
	// constant: ConstB
	// constant: ConstC
}
