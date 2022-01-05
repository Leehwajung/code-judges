package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	util.CheckArgsMinCount(1)
	s := os.Args[1]
	output := lengthOfLongestSubstring(s)
	fmt.Println(output)
}
