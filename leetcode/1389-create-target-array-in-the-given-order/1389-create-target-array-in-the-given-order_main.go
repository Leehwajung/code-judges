package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
	"code-judges/internal/typeconv"
)

func main() {
	args.MustEnough(2)
	nums := typeconv.StringsToInts(os.Args[1 : len(os.Args)/2+1])
	index := typeconv.StringsToInts(os.Args[len(os.Args)/2+1:])
	output := createTargetArray(nums, index)
	fmt.Print(output)
}
