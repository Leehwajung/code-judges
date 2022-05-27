package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
	"code-judges/internal/typeconv"
)

func main() {
	args.MustEnough(1)
	n := typeconv.StringToInt(os.Args[1])
	output := fizzBuzz(n)
	fmt.Print(output)
}
