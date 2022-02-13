package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	arr := util.StringsToInts(os.Args[1:])
	output := canMakeArithmeticProgression(arr)
	fmt.Println(output)
}
