package main

import (
	"fmt"
	"os"

	"code-judges/internal/typeconv"
)

func main() {
	nums := typeconv.StringsToInts(os.Args[1:])
	output := findMin(nums)
	fmt.Print(output)
}
