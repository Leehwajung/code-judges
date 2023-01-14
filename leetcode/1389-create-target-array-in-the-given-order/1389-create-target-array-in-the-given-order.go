package main

func createTargetArray(nums []int, index []int) []int {
	var result []int
	for i, idx := range index {
		result = append(result[:idx], append([]int{nums[i]}, result[idx:]...)...)
	}
	return result
}
