package main

import (
	"fmt"
)

func intOrNilsToBinaryTree(vals ...interface{}) *TreeNode {
	return intOrNilsToBinaryTreeRecursive(vals, 0)
}

func intOrNilsToBinaryTreeRecursive(vals []interface{}, index int) *TreeNode {
	if index >= len(vals) || vals[index] == nil {
		return nil
	}
	return &TreeNode{
		Val:   vals[index].(int),
		Left:  intOrNilsToBinaryTreeRecursive(vals, index*2+1),
		Right: intOrNilsToBinaryTreeRecursive(vals, index*2+2),
	}
}

func binaryTreeToIntOrNils(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}

	valMap := make(map[int]int)
	binaryTreeToIntOrNilsRecursive(root, valMap, 0)

	maxIdx := 0
	for idx := range valMap {
		if maxIdx < idx {
			maxIdx = idx
		}
	}

	valsLen := 0
	levelBegin := 1
	for valsLen <= maxIdx {
		valsLen += levelBegin
		levelBegin *= 2
	}

	vals := make([]interface{}, valsLen)
	for idx, val := range valMap {
		vals[idx] = val
	}
	return vals
}

func binaryTreeToIntOrNilsRecursive(node *TreeNode, outVals map[int]int, index int) {
	if node == nil {
		return
	}
	outVals[index] = node.Val
	binaryTreeToIntOrNilsRecursive(node.Left, outVals, index*2+1)
	binaryTreeToIntOrNilsRecursive(node.Right, outVals, index*2+2)
}

func (t *TreeNode) String() string {
	return fmt.Sprint(binaryTreeToIntOrNils(t))
}
