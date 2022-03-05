package main

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	expandableCases := make([]int, 0, len(nums)/2)
	subarrays := 0
	cases := 1

	for idx := 0; idx <= len(nums); idx++ {
		if idx < len(nums) && nums[idx] <= right {
			if nums[idx] < left {
				cases++
			} else {
				expandableCases = append(expandableCases, cases)
				cases = 1
			}
			continue
		}
		expandableCases = append(expandableCases, cases)

		if len(expandableCases) > 1 {
			sum := expandableCases[0]
			for _, expandableCase := range expandableCases[1:] {
				subarrays += sum * expandableCase
				sum += expandableCase
			}
		}

		cases = 1
		expandableCases = expandableCases[0:0]
	}

	return subarrays
}
