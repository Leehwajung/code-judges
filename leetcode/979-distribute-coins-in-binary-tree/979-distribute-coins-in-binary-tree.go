package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	moves := 0
	_ = getDistributedVal(root, &moves)
	return moves
}

func getDistributedVal(node *TreeNode, moves *int) (newVal int) {
	newVal = node.Val - 1
	if node.Left != nil {
		leftVal := getDistributedVal(node.Left, moves)
		*moves += abs(leftVal)
		newVal += leftVal
	}
	if node.Right != nil {
		rightVal := getDistributedVal(node.Right, moves)
		*moves += abs(rightVal)
		newVal += rightVal
	}
	return newVal
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
