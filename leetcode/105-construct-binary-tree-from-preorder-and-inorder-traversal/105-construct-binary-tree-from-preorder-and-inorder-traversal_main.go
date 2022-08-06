package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
	"code-judges/internal/typeconv"
)

func main() {
	args.MustEnough(2)
	args.MustEvenCount()
	a := os.Args[1:]
	preorder := typeconv.StringsToInts(a[:len(a)/2])
	inorder := typeconv.StringsToInts(a[len(a)/2:])
	tree := buildTree(preorder, inorder)
	fmt.Print(tree)
}
