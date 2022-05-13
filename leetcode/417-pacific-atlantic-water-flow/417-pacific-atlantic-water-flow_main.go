package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	util.CheckArgsMinCount(2)
	rowLen := util.StringToInt(os.Args[len(os.Args)-1])
	heights := util.StringsToIntGrid(os.Args[1:len(os.Args)-1], rowLen)
	output := pacificAtlantic(heights)
	fmt.Println(output)
}
