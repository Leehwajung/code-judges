package main

import (
	"fmt"
	"os"

	"code-judges/internal/typeconv"
)

func main() {
	arr := typeconv.StringsToInts(os.Args[1:])
	output := canMakeArithmeticProgression(arr)
	fmt.Print(output)
}
