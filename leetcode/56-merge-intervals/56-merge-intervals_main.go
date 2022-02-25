package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	util.CheckArgsMinCount(2)
	intervals := util.StringsToIntPairs(os.Args[1:])
	output := merge(intervals)
	fmt.Println(output)
}
