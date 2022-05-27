package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
	"code-judges/internal/typeconv"
)

func main() {
	args.MustEnough(3)
	nums := typeconv.StringsToInts(os.Args[1 : len(os.Args)-2])
	left := typeconv.StringToInt(os.Args[len(os.Args)-2])
	right := typeconv.StringToInt(os.Args[len(os.Args)-1])
	output := numSubarrayBoundedMax(nums, left, right)
	fmt.Print(output)
}
