package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
	"code-judges/internal/typeconv"
)

func main() {
	args.MustEnough(2)
	height := typeconv.StringsToInts(os.Args[1:])
	output := maxArea(height)
	fmt.Print(output)
}
