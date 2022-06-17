package main

func numEnclaves(grid [][]int) int {
	for nIdx := 0; nIdx < len(grid[0]); nIdx++ {
		walk(grid, 0, nIdx)
	}
	for mIdx := 1; mIdx < len(grid)-1; mIdx++ {
		walk(grid, mIdx, 0)
		walk(grid, mIdx, len(grid[mIdx])-1)
	}
	for nIdx := 0; nIdx < len(grid[len(grid)-1]); nIdx++ {
		walk(grid, len(grid)-1, nIdx)
	}

	enclaveCellCount := 0
	for mIdx := 1; mIdx < len(grid)-1; mIdx++ {
		for nIdx := 1; nIdx < len(grid[mIdx])-1; nIdx++ {
			if grid[mIdx][nIdx] != 0 {
				enclaveCellCount++
			}
		}
	}
	return enclaveCellCount
}

func walk(grid [][]int, mIdx, nIdx int) {
	if mIdx < 0 || mIdx >= len(grid) {
		return
	}
	if nIdx < 0 || nIdx >= len(grid[mIdx]) {
		return
	}
	if grid[mIdx][nIdx] == 0 {
		return
	}

	grid[mIdx][nIdx] = 0
	walk(grid, mIdx-1, nIdx)
	walk(grid, mIdx, nIdx-1)
	walk(grid, mIdx+1, nIdx)
	walk(grid, mIdx, nIdx+1)
}
