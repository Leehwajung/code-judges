package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	util.CheckArgsMinCount(1)
	num := util.BitStringToUint32(os.Args[1])
	output := reverseBits(num)
	fmt.Printf("%b\n", output)
}
