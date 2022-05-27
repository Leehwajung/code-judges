package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
	"code-judges/internal/typeconv"
)

func main() {
	args.MustEnough(1)
	num := typeconv.BitStringToUint32(os.Args[1])
	output := reverseBits(num)
	fmt.Printf("%b\n", output)
}
