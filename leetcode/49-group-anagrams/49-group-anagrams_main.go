package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	util.CheckArgsMinCount(1)
	strs := os.Args[1:]
	output := groupAnagrams(strs)
	fmt.Println(output)
}
