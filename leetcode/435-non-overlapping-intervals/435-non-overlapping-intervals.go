package main

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1]
	})

	removes := 0
	prev := intervals[0]
	for _, curr := range intervals[1:] {
		if curr[0] < prev[1] {
			if curr[1] < prev[1] {
				prev = curr
			}
			removes++
		} else {
			prev = curr
		}
	}
	return removes
}
