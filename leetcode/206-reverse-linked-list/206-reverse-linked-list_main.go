package main

import (
	"fmt"
	"os"

	"code-judges/internal/util"
)

func main() {
	util.CheckArgsMinCount(2)
	head := intsToLinkedList(util.StringsToInts(os.Args[1:])...)
	output := reverseList(head)
	fmt.Println(linkedListToInts(output))
}
