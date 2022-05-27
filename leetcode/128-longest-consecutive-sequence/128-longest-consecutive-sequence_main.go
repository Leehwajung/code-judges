package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
	"code-judges/internal/typeconv"
)

func main() {
	args.MustEnough(0)
	nums := typeconv.StringsToInts(os.Args[1:])
	output := longestConsecutive(nums)
	fmt.Print(output)
}
