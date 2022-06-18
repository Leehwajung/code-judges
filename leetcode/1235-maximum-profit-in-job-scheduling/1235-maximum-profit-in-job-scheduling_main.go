package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
	"code-judges/internal/typeconv"
)

func main() {
	args.MustEnough(3)
	parsedArgs := typeconv.StringsToIntGrid(os.Args[1:], len(os.Args[1:])/3)
	startTime := parsedArgs[0]
	endTime := parsedArgs[1]
	profit := parsedArgs[2]
	output := jobScheduling(startTime, endTime, profit)
	fmt.Print(output)
}
