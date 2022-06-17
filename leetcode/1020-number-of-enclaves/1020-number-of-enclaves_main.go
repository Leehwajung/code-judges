package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
	"code-judges/internal/typeconv"
)

func main() {
	args.MustEnough(2)
	rowLen := typeconv.StringToInt(os.Args[len(os.Args)-1])
	grid := typeconv.StringsToIntGrid(os.Args[1:len(os.Args)-1], rowLen)
	output := numEnclaves(grid)
	fmt.Print(output)
}
