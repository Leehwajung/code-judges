package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	nums := util.StringsToInts(os.Args[1 : len(os.Args)-1])
	target := util.StringToInt(os.Args[len(os.Args)-1])
	output := search(nums, target)
	fmt.Println(output)
}
