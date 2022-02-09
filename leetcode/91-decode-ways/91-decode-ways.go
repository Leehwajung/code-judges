package main

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	nums := make([]int, len(s))
	nums[0] = 1

	for idx := 1; idx < len(s); idx++ {
		prev1Idx := idx - 1
		prev1Num := nums[prev1Idx]
		if s[idx] == '0' {
			if s[prev1Idx] != '1' && s[prev1Idx] != '2' {
				return 0
			}
			prev1Num = 0
		}

		prev2Num := 0
		if s[prev1Idx] == '1' || (s[prev1Idx] == '2' && s[idx] <= '6') {
			if prev1Idx > 0 {
				prev2Num = nums[prev1Idx-1]
			} else {
				prev2Num = 1
			}
		}

		nums[idx] = prev1Num + prev2Num
	}

	return nums[len(s)-1]
}
