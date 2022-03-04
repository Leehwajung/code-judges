package main

func gridGame(grid [][]int) int64 {
	sums := make([]int, len(grid[0]))
	row0Sum := 0
	for i := len(grid[0]) - 1; i >= 0; i-- {
		row0Sum += grid[0][i]
		sums[i] = row0Sum
	}

	maxSum := 0
	maxIdx := 0
	row1Sum := 0
	for i := 0; i < len(grid[1]); i++ {
		row1Sum += grid[1][i]
		if row1Sum < sums[i] {
			sums[i] = row1Sum
		}
		if sums[i] > maxSum {
			maxSum = sums[i]
			maxIdx = i
		}
	}

	switch {
	case len(sums) <= 1:
		return 0
	case maxIdx == 0:
		return int64(sums[maxIdx+1])
	case maxIdx == len(sums)-1:
		return int64(sums[maxIdx-1])
	default:
		if sums[maxIdx+1] > sums[maxIdx-1] {
			return int64(sums[maxIdx+1])
		}
		return int64(sums[maxIdx-1])
	}
}
