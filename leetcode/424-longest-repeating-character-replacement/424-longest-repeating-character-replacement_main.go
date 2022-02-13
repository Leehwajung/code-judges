package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	util.CheckArgsMinCount(2)
	s := os.Args[1]
	k := util.StringToInt(os.Args[2])
	output := characterReplacement(s, k)
	fmt.Println(output)
}
