package main

func checkSubarraySum(nums []int, k int) bool {
	indexes := map[int]int{0: -1}
	sumRem := 0
	for idx, num := range nums {
		sumRem = (sumRem + num) % k
		if sumIdx, ok := indexes[sumRem]; !ok {
			indexes[sumRem] = idx
		} else if idx-sumIdx >= 2 {
			return true
		}
	}
	return false
}
