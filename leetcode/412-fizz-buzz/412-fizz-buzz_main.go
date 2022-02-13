package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	util.CheckArgsMinCount(1)
	n := util.StringToInt(os.Args[1])
	output := fizzBuzz(n)
	fmt.Println(output)
}
