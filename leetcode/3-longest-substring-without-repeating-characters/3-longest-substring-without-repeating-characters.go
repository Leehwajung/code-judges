package main

func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	if sLen <= 1 {
		return sLen
	}

	table := make([]bool, 127)
	table[s[0]] = true
	longest, begin := 1, 0
	for idx := 1; idx < sLen; idx++ {
		if table[s[idx]] {
			for s[begin] != s[idx] {
				table[s[begin]] = false
				begin++
			}
			begin++
		} else if longest <= idx-begin {
			longest = idx - begin + 1
		}
		table[s[idx]] = true
	}
	return longest
}
