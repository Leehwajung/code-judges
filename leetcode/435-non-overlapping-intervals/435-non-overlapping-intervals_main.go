package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
	"code-judges/internal/typeconv"
)

func main() {
	args.MustEnough(2)
	intervals := typeconv.StringsToIntPairs(os.Args[1:])
	output := eraseOverlapIntervals(intervals)
	fmt.Print(output)
}
