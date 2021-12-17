package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	util.CheckArgsMinCount(1)
	coins := util.StringsToInts(os.Args[1 : len(os.Args)-1])
	amount := util.StringToInt(os.Args[len(os.Args)-1])
	output := coinChange(coins, amount)
	fmt.Println(output)
}
