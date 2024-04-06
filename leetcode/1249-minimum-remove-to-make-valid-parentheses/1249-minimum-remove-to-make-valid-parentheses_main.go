package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
)

func main() {
	args.MustEnough(1)
	output := minRemoveToMakeValid(os.Args[1])
	fmt.Print(output)
}
