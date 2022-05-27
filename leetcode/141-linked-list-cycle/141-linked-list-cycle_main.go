package main

import (
	"fmt"
	"os"

	"code-judges/internal/typeconv"
)

func main() {
	list := []int(nil)
	pos := -1
	if len(os.Args)-1 > 0 {
		list = typeconv.StringsToInts(os.Args[1 : len(os.Args)-1])
		pos = typeconv.StringToInt(os.Args[len(os.Args)-1])
	}
	head := linkedListWithCycle(pos, list...)
	output := hasCycle(head)
	fmt.Print(output)
}
