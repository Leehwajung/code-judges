package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
)

func main() {
	args.MustEnough(2)
	output := wordBreak(os.Args[1], os.Args[2:])
	fmt.Print(output)
}
