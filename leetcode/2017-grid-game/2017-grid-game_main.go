package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	util.CheckArgsMinCount(2)
	util.MustBeEvenArguments()
	grid := [][]int{
		util.StringsToInts(os.Args[1 : len(os.Args)/2+1]),
		util.StringsToInts(os.Args[len(os.Args)/2+1:]),
	}
	output := gridGame(grid)
	fmt.Println(output)
}
