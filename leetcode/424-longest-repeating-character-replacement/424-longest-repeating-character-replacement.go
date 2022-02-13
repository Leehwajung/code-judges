package main

import (
	"fmt"
)

func characterReplacement7(s string, k int) int {
	wholeLen := len(s)
	prevLetter := byte(0)
	chanceTable := make(map[int]int)
	result := 0

	for currIdx := 0; currIdx < wholeLen; currIdx++ {
		for beginIdx, chances := range chanceTable {
			if s[beginIdx] == s[currIdx] {
				continue
			} else if chances > 0 {
				chanceTable[beginIdx]--
				continue
			}
			subLen := currIdx - beginIdx
			if result < subLen {
				result = subLen
			}
			delete(chanceTable, beginIdx)
		}

		if s[currIdx] != prevLetter {
			chanceTable[currIdx] = k
		}
		prevLetter = s[currIdx]
	}

	for beginIdx, chances := range chanceTable {
		subLen := wholeLen - beginIdx + chances
		if subLen >= wholeLen {
			result = wholeLen
			break
		}
		if result < subLen {
			result = subLen
		}
	}
	return result
}

func characterReplacement6(s string, k int) int {
	wholeLen := len(s)
	prevLetter := byte(0)
	chanceTable := make(map[int]int)
	result := 0

	for currIdx := 0; currIdx < wholeLen; currIdx++ {
		shouldPut := s[currIdx] != prevLetter
		prevLetter = s[currIdx]

		for beginIdx, chances := range chanceTable {
			if s[beginIdx] == s[currIdx] {
				shouldPut = false
				continue
			} else if chances > 0 {
				chanceTable[beginIdx]--
				continue
			}
			subLen := currIdx - beginIdx
			if result < subLen {
				result = subLen
			}
			delete(chanceTable, beginIdx)
		}

		if shouldPut {
			chanceTable[currIdx] = k
		}
	}

	for beginIdx, chances := range chanceTable {
		subLen := wholeLen - beginIdx + chances
		if subLen >= wholeLen {
			result = wholeLen
			break
		}
		if result < subLen {
			result = subLen
		}
	}
	return result
}

func characterReplacement5(s string, k int) int {
	wholeLen := len(s)
	prevLetter := byte(0)
	chanceTable := make(map[int]int)
	result := 0

	for currIdx := 0; currIdx < wholeLen; currIdx++ {
		for beginIdx, chances := range chanceTable {
			if s[beginIdx] == s[currIdx] {
				continue
			} else if chances > 0 {
				chanceTable[beginIdx]--
				continue
			}
			subLen := currIdx - beginIdx
			if result < subLen {
				result = subLen
			}
			delete(chanceTable, beginIdx)
		}

		if s[currIdx] != prevLetter {
			chanceTable[currIdx] = k
		}
		prevLetter = s[currIdx]
	}

	for beginIdx, chances := range chanceTable {
		if chances > 0 {
			currIdx := beginIdx
			for currIdx > 0 && chances > 0 {
				currIdx--
				if s[currIdx] != s[beginIdx] {
					chances--
				}
			}
			beginIdx = currIdx
		}

		subLen := wholeLen - beginIdx
		if result < subLen {
			result = subLen
		}
	}
	return result
}

func characterReplacement4(s string, k int) int {
	wholeLen := len(s)
	//prevLetter := byte(0)
	chanceTable := make(map[int]int)
	result := 0

	for currIdx := 0; currIdx < wholeLen; currIdx++ {
		//same := s[currIdx] == prevLetter
		//prevLetter = s[currIdx]
		//if same {
		//	continue
		//}

		for beginIdx, chances := range chanceTable {
			if s[beginIdx] == s[currIdx] {
				continue
			} else if chances > 0 {
				chanceTable[beginIdx]--
				continue
			}
			subLen := currIdx - beginIdx
			if result < subLen {
				result = subLen
			}
			delete(chanceTable, beginIdx)
		}
		chanceTable[currIdx] = k
	}

	for beginIdx := range chanceTable {
		subLen := wholeLen - beginIdx
		if result < subLen {
			result = subLen
		}
	}
	return result
}

func characterReplacement3(s string, k int) int {
	wholeLen := len(s)
	prevLetter := byte(0)
	chanceTable := make(map[int]int)
	result := 0

	for currIdx := 0; currIdx < wholeLen; currIdx++ {
		same := s[currIdx] == prevLetter
		prevLetter = s[currIdx]
		if same {
			continue
		}

		for beginIdx, chances := range chanceTable {
			if chances > 0 {
				chanceTable[beginIdx]--
				continue
			}
			subLen := (currIdx + 1) - beginIdx
			if result < subLen {
				result = subLen
			}
			delete(chanceTable, beginIdx)
		}
		chanceTable[currIdx] = k
	}

	for beginIdx := range chanceTable {
		subLen := wholeLen - beginIdx
		if result < subLen {
			result = subLen
		}
	}
	return result
}

func characterReplacement2(s string, k int) int {
	letter := byte(0)
	begin := 0

	table := make([]int, len(s)+1)
	for curLen := 1; curLen <= len(s); curLen++ {
		// k = 1, idx = 1, len = 2
		// k = 2, idx = 2, len = 3
		idx := curLen - 1
		if k >= idx {
			table[curLen] = curLen
			continue
		}
		if letter == s[idx] {

		} else {
			begin = idx
		}
	}
	fmt.Println(begin)
	return table[len(s)]
}

func characterReplacement1(s string, k int) int {
	if len(s) <= 1 {
		return len(s)
	}
	table := make([]int, len(s)+1)
	for curLen := 1; curLen <= len(s); curLen++ {
		// k = 1, idx = 1, len = 2
		// k = 2, idx = 2, len = 3
		idx := curLen - 1
		if k >= idx {
			table[curLen] = curLen
		} else {

		}
	}
	return table[len(s)]
}
