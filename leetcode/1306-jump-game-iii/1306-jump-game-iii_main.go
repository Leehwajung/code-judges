package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
	"code-judges/internal/typeconv"
)

func main() {
	args.MustEnough(2)
	arr := typeconv.StringsToInts(os.Args[1 : len(os.Args)-1])
	start := typeconv.StringToInt(os.Args[len(os.Args)-1])
	output := canReach(arr, start)
	fmt.Print(output)
}
