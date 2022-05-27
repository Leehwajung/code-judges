package main

func longestConsecutive(nums []int) int {
	longestLen := 0
	numSet := make(map[int]bool)

	for _, num := range nums {
		numSet[num] = true
	}

	for num := range numSet {
		currLen := 1
		delete(numSet, num)

		for afgerNum := num + 1; numSet[afgerNum]; afgerNum++ {
			currLen++
			delete(numSet, afgerNum)
		}

		for beforeNum := num - 1; numSet[beforeNum]; beforeNum-- {
			currLen++
			delete(numSet, beforeNum)
		}

		if longestLen < currLen {
			longestLen = currLen
		}
	}

	return longestLen
}
