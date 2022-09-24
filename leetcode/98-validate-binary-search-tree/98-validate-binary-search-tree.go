package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	validator := bstValidator{prevVal: math.MinInt}
	return validator.validate(root)
}

type bstValidator struct {
	prevVal int
}

func (v *bstValidator) validate(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !v.validate(root.Left) {
		return false
	}

	if root.Val <= v.prevVal {
		return false
	}
	v.prevVal = root.Val

	if !v.validate(root.Right) {
		return false
	}

	return true
}
