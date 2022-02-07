package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	nums := util.StringsToInts(os.Args[1:])
	output := findMin(nums)
	fmt.Println(output)
}
