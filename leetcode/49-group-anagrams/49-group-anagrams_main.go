package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
)

func main() {
	args.MustEnough(1)
	strs := os.Args[1:]
	output := groupAnagrams(strs)
	fmt.Print(output)
}
