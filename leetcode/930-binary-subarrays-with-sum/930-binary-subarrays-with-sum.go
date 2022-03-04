package main

func numSubarraysWithSum(nums []int, goal int) int {
	zeroContCounts := make([]int, 0, len(nums))
	zeroContCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroContCount++
		} else {
			zeroContCounts = append(zeroContCounts, zeroContCount)
			zeroContCount = 0
		}
	}
	zeroContCounts = append(zeroContCounts, zeroContCount)

	subarrays := 0
	if goal == 0 {
		for _, zeroContCount := range zeroContCounts {
			subarrays += zeroContCount * (zeroContCount + 1) / 2
		}
	} else {
		for i := 0; i < len(zeroContCounts)-goal; i++ {
			subarrays += (zeroContCounts[i] + 1) * (zeroContCounts[i+goal] + 1)
		}
	}
	return subarrays
}
