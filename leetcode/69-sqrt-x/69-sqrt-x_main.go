package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
	"code-judges/internal/typeconv"
)

func main() {
	args.MustEnough(1)
	x := typeconv.StringToInt(os.Args[1])
	output := mySqrt(x)
	fmt.Print(output)
}
