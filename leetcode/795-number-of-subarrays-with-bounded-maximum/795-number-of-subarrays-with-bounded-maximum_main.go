package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	util.CheckArgsMinCount(3)
	nums := util.StringsToInts(os.Args[1 : len(os.Args)-2])
	left := util.StringToInt(os.Args[len(os.Args)-2])
	right := util.StringToInt(os.Args[len(os.Args)-1])
	output := numSubarrayBoundedMax(nums, left, right)
	fmt.Println(output)
}
