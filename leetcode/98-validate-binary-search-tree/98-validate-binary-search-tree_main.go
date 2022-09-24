package main

import (
	"fmt"
	"os"

	"code-judges/internal/typeconv"
)

func main() {
	intOrNils := typeconv.StringsToIntOrNils(os.Args[1:])
	binaryTree := intOrNilsToBinaryTree(intOrNils...)
	output := isValidBST(binaryTree)
	fmt.Print(output)
}
