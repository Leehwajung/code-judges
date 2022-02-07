package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	util.CheckArgsMinCount(1)
	x := util.StringToInt(os.Args[1])
	output := mySqrt(x)
	fmt.Println(output)
}
