package main

func pacificAtlantic(heights [][]int) (coordinates [][]int) {
	for r := 0; r < len(heights); r++ {
		row := heights[r]
		for c := 0; c < len(row); c++ {
			if canFlowToPacific(heights, r, c) && canFlowToAtlantic(heights, r, c) {
				coordinates = append(coordinates, []int{r, c})
			}
		}
	}
	return coordinates
}

func canFlowToPacific(heights [][]int, r, c int) bool {
	visits := make(map[struct{ r, c int }]bool)
	flowPredicate := func(r, c int) bool { return r == 0 || c == 0 }
	return canFlowRecursive(heights, r, c, visits, flowPredicate)
}

func canFlowToAtlantic(heights [][]int, r, c int) bool {
	visits := make(map[struct{ r, c int }]bool)
	flowPredicate := func(r, c int) bool { return r == len(heights)-1 || c == len(heights[0])-1 }
	return canFlowRecursive(heights, r, c, visits, flowPredicate)
}

func canFlowRecursive(heights [][]int, r, c int, visits map[struct{ r, c int }]bool, flowPredicate func(r, c int) bool) bool {
	if flowPredicate(r, c) {
		return true
	} else if visits[struct{ r, c int }{r: r, c: c}] {
		return false
	}

	visits[struct{ r, c int }{r: r, c: c}] = true
	if r > 0 && heights[r][c] >= heights[r-1][c] && canFlowRecursive(heights, r-1, c, visits, flowPredicate) {
		return true
	} else if c > 0 && heights[r][c] >= heights[r][c-1] && canFlowRecursive(heights, r, c-1, visits, flowPredicate) {
		return true
	} else if r < len(heights)-1 && heights[r][c] >= heights[r+1][c] && canFlowRecursive(heights, r+1, c, visits, flowPredicate) {
		return true
	} else if c < len(heights[0])-1 && heights[r][c] >= heights[r][c+1] && canFlowRecursive(heights, r, c+1, visits, flowPredicate) {
		return true
	}
	return false
}
