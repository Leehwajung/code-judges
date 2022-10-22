package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	var results []float64
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelNodeValTotal := 0
		levelNodeCount := len(queue)

		for i := 0; i < levelNodeCount; i++ {
			front := queue[0]
			queue = queue[1:]

			levelNodeValTotal += front.Val

			if front.Left != nil {
				queue = append(queue, front.Left)
			}

			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		}

		levelAverage := float64(levelNodeValTotal) / float64(levelNodeCount)
		results = append(results, levelAverage)
	}

	return results
}
