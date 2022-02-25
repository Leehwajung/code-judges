package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return [][]int{{intervals[0][0], intervals[0][1]}}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	mInt := []int{intervals[0][0], intervals[0][1]}
	result := make([][]int, 0, len(intervals)/2)
	for idx := 1; idx < len(intervals); idx++ {
		if intervals[idx][0] < mInt[0] || intervals[idx][0] > mInt[1] {
			result = append(result, mInt)
			mInt = []int{intervals[idx][0], intervals[idx][1]}
		} else if intervals[idx][1] > mInt[1] {
			mInt[1] = intervals[idx][1]
		}
	}
	return append(result, mInt)
}
