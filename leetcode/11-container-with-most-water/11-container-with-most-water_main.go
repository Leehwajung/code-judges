package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	util.CheckArgsMinCount(2)
	height := util.StringsToInts(os.Args[1:])
	output := maxArea(height)
	fmt.Println(output)
}
