package main

import (
	"fmt"
	"os"

	"code-judges/internal/args"
	"code-judges/internal/typeconv"
)

func main() {
	args.MustEnough(2)
	head := intsToLinkedList(typeconv.StringsToInts(os.Args[1:])...)
	output := reverseList(head)
	fmt.Println(linkedListToInts(output))
}
