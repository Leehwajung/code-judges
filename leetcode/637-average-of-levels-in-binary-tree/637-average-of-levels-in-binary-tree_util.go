package main

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
