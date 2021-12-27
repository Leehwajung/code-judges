package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	list := []int(nil)
	pos := -1
	if len(os.Args)-1 > 0 {
		list = util.StringsToInts(os.Args[1 : len(os.Args)-1])
		pos = util.StringToInt(os.Args[len(os.Args)-1])
	}
	head := linkedListWithCycle(pos, list...)
	output := hasCycle(head)
	fmt.Println(output)
}
