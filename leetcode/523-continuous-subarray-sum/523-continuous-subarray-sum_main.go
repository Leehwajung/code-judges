package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	util.CheckArgsMinCount(2)
	nums := util.StringsToInts(os.Args[1 : len(os.Args)-1])
	k := util.StringToInt(os.Args[len(os.Args)-1])
	output := checkSubarraySum(nums, k)
	fmt.Println(output)
}
