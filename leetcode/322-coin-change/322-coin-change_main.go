package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
	"code-judges/internal/typeconv"
)

func main() {
	args.MustEnough(1)
	coins := typeconv.StringsToInts(os.Args[1 : len(os.Args)-1])
	amount := typeconv.StringToInt(os.Args[len(os.Args)-1])
	output := coinChange(coins, amount)
	fmt.Print(output)
}
