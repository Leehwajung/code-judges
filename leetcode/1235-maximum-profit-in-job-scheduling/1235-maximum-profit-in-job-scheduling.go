package main

import (
	"sort"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	sort.Sort(jobs{startTime: startTime, endTime: endTime, profit: profit})

	maxProfit := make([]int, len(profit))
	maxProfit[0] = profit[0]

	for idx := 1; idx < len(profit); idx++ {
		prevMaxIdx := sort.SearchInts(endTime, startTime[idx])
		prevMax := 0
		if endTime[prevMaxIdx] == startTime[idx] {
			prevMaxIdx = sort.SearchInts(endTime, startTime[idx]+1) - 1
			prevMax = maxProfit[prevMaxIdx]
		} else if prevMaxIdx > 0 {
			prevMax = maxProfit[prevMaxIdx-1]
		}

		maxProfit[idx] = prevMax + profit[idx]
		if maxProfit[idx] < maxProfit[idx-1] {
			maxProfit[idx] = maxProfit[idx-1]
		}
	}

	return maxProfit[len(maxProfit)-1]
}

type jobs struct{ startTime, endTime, profit []int }

func (js jobs) Len() int           { return len(js.endTime) }
func (js jobs) Less(i, j int) bool { return js.endTime[i] < js.endTime[j] }
func (js jobs) Swap(i, j int) {
	js.startTime[i], js.startTime[j] = js.startTime[j], js.startTime[i]
	js.endTime[i], js.endTime[j] = js.endTime[j], js.endTime[i]
	js.profit[i], js.profit[j] = js.profit[j], js.profit[i]
}
