package main

func uniquePaths(m int, n int) int {
	table := make([][]int, m)
	{
		table[0] = make([]int, n)
		for nIdx := 0; nIdx < n; nIdx++ {
			table[0][nIdx] = 1
		}
	}
	for mIdx := 1; mIdx < m; mIdx++ {
		table[mIdx] = make([]int, n)
		table[mIdx][0] = 1
		for nIdx := 1; nIdx < n; nIdx++ {
			table[mIdx][nIdx] = table[mIdx-1][nIdx] + table[mIdx][nIdx-1]
		}
	}
	return table[m-1][n-1]
}
