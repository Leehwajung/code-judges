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
	intervals := typeconv.StringsToIntPairs(os.Args[1:])
	output := merge(intervals)
	fmt.Print(output)
}
