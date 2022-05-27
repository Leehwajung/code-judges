package main

import (
	"fmt"
	"os"

	"code-judges/internal/typeconv"
)

func main() {
	nums := typeconv.StringsToInts(os.Args[1 : len(os.Args)-1])
	target := typeconv.StringToInt(os.Args[len(os.Args)-1])
	output := search(nums, target)
	fmt.Print(output)
}
