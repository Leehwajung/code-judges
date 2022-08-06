package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	inorderRootIdx := -1
	for idx, val := range inorder {
		if val == rootVal {
			inorderRootIdx = idx
			break
		}
	}
	var left, right *TreeNode
	if 1 <= inorderRootIdx && inorderRootIdx < len(preorder) {
		left = buildTree(preorder[1:inorderRootIdx+1], inorder[:inorderRootIdx])
	}
	if 0 <= inorderRootIdx && inorderRootIdx < len(preorder)-1 {
		right = buildTree(preorder[inorderRootIdx+1:], inorder[inorderRootIdx+1:])
	}
	return &TreeNode{
		Val:   rootVal,
		Left:  left,
		Right: right,
	}
}
