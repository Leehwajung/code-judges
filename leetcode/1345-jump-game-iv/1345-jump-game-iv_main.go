package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
	"code-judges/internal/typeconv"
)

func main() {
	args.MustEnough(1)
	arr := typeconv.StringsToInts(os.Args[1:])
	output := minJumps(arr)
	fmt.Print(output)
}
