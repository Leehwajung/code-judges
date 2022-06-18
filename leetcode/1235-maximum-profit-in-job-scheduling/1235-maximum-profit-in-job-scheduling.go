package main

import (
	"sort"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	sort.Sort(jobs{startTime: startTime, endTime: endTime, profit: profit})

	maxProfit := make([]int, len(profit))
	maxProfit[len(maxProfit)-1] = profit[len(profit)-1]
	for idx := len(profit) - 1; idx >= 0; idx-- {
		nextStart := sort.SearchInts(startTime, endTime[idx])
		maxProfit[idx] = profit[idx]
		// TODO 같은 start 시간에서 최대 profit 찾기
		if nextStart < len(startTime) {
			maxProfit[idx] += maxProfit[nextStart]
		}
	}
	return maxProfit[0]
}

type jobs struct{ startTime, endTime, profit []int }

func (js jobs) Len() int { return len(js.startTime) }

func (js jobs) Less(i, j int) bool {
	if js.startTime[i] != js.startTime[j] {
		return js.startTime[i] < js.startTime[j]
	} else {
		return js.profit[i] > js.profit[j]
	}
}

func (js jobs) Swap(i, j int) {
	js.startTime[i], js.startTime[j] = js.startTime[j], js.startTime[i]
	js.endTime[i], js.endTime[j] = js.endTime[j], js.endTime[i]
	js.profit[i], js.profit[j] = js.profit[j], js.profit[i]
}
