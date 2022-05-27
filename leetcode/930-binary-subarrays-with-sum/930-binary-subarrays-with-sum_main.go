package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
	"code-judges/internal/typeconv"
)

func main() {
	args.MustEnough(2)
	nums := typeconv.StringsToInts(os.Args[1 : len(os.Args)-1])
	goal := typeconv.StringToInt(os.Args[len(os.Args)-1])
	output := numSubarraysWithSum(nums, goal)
	fmt.Print(output)
}
