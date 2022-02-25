package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	util.CheckArgsMinCount(2)
	intervals := util.StringsToIntPairs(os.Args[1 : len(os.Args)-2])
	newInterval := util.StringsToInts(os.Args[len(os.Args)-2:])
	output := insert(intervals, newInterval)
	fmt.Println(output)
}
