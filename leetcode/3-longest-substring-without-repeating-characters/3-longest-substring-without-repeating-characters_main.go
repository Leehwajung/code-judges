package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
)

func main() {
	args.MustEnough(1)
	s := os.Args[1]
	output := lengthOfLongestSubstring(s)
	fmt.Print(output)
}
