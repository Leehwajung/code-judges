package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
	"code-judges/internal/typeconv"
)

func main() {
	args.MustEnough(2)
	m := typeconv.StringToInt(os.Args[1])
	n := typeconv.StringToInt(os.Args[2])
	output := uniquePaths(m, n)
	fmt.Print(output)
}
