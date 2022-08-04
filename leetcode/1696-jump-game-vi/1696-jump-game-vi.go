package main

func maxResult(nums []int, k int) int {
	idxDeque := []int{0}

	for idx := 1; idx < len(nums); idx++ {
		if idxDeque[0]+k < idx {
			idxDeque = idxDeque[1:]
		}

		nums[idx] += nums[idxDeque[0]]

		for len(idxDeque) > 0 && nums[idxDeque[len(idxDeque)-1]] <= nums[idx] {
			idxDeque = idxDeque[:len(idxDeque)-1]
		}

		idxDeque = append(idxDeque, idx)
	}

	return nums[idxDeque[len(idxDeque)-1]]
}
