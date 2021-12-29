package main

import (
	"strings"
)

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	} else if text1 == text2 {
		return len(text1)
	} else if strings.Contains(text1, text2) {
		return len(text2)
	} else if strings.Contains(text2, text1) {
		return len(text1)
	}

	table := make([][]int, len(text1))
	result := 0

	{
		table[0] = make([]int, len(text2))
		if text1[0] == text2[0] {
			table[0][0] = 1
			result = 1
		}

		for j := 1; j < len(text2); j++ {
			if text1[0] == text2[j] {
				table[0][j] = 1
			} else {
				table[0][j] = table[0][j-1]
			}
			if result < table[0][j] {
				result = table[0][j]
			}
		}
	}

	for i := 1; i < len(text1); i++ {
		table[i] = make([]int, len(text2))
		{
			if text1[i] == text2[0] {
				table[i][0] = 1
			} else {
				table[i][0] = table[i-1][0]
			}
			if result < table[i][0] {
				result = table[i][0]
			}
		}

		for j := 1; j < len(text2); j++ {
			if text1[i] == text2[j] {
				table[i][j] = table[i-1][j-1] + 1
			} else if table[i-1][j] > table[i][j-1] {
				table[i][j] = table[i-1][j]
			} else {
				table[i][j] = table[i][j-1]
			}
			if result < table[i][j] {
				result = table[i][j]
			}
		}
	}

	return result
}
