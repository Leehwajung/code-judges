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
	intervals := typeconv.StringsToIntPairs(os.Args[1 : len(os.Args)-2])
	newInterval := typeconv.StringsToInts(os.Args[len(os.Args)-2:])
	output := insert(intervals, newInterval)
	fmt.Print(output)
}
