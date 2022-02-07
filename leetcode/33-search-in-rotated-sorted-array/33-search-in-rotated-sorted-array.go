package main

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[left] == target {
			return left
		} else if nums[right] == target {
			return right
		}

		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < nums[right] {
			if nums[mid] < target && target < nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[right] < target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}
