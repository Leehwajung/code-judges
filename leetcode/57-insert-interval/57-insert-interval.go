package main

import (
	"sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	idx := sort.Search(len(intervals), func(i int) bool {
		return intervals[i][0] >= newInterval[0]
	})

	oldIntervals := intervals
	intervals = make([][]int, len(oldIntervals)+1)
	for i := 0; i < idx; i++ { // To avoid a side effect for argument
		intervals[i] = []int{oldIntervals[i][0], oldIntervals[i][1]}
	}
	intervals[idx] = newInterval
	copy(intervals[idx+1:], oldIntervals[idx:])

	if idx > 0 && (idx >= len(oldIntervals) || newInterval[0] != oldIntervals[idx][0]) {
		idx--
	}

	mInt := []int{intervals[idx][0], intervals[idx][1]}
	result := intervals[:idx]
	for idx = idx + 1; idx < len(intervals); idx++ {
		if intervals[idx][0] < mInt[0] || intervals[idx][0] > mInt[1] {
			result = append(result, mInt)
			mInt = []int{intervals[idx][0], intervals[idx][1]}
		} else if intervals[idx][1] > mInt[1] {
			mInt[1] = intervals[idx][1]
		}
	}
	return append(result, mInt)
}
