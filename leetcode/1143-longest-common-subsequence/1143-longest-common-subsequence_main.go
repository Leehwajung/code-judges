package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	util.CheckArgsMinCount(2)
	text1 := os.Args[1]
	text2 := os.Args[2]
	output := longestCommonSubsequence(text1, text2)
	fmt.Println(output)
}
