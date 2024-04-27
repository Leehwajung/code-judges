package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
)

func main() {
	args.MustEnough(2)
	deadends := os.Args[1 : len(os.Args)-1]
	target := os.Args[len(os.Args)-1]
	output := openLock(deadends, target)
	fmt.Print(output)
}
