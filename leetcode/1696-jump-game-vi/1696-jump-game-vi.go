package main

import (
	"sort"
)

func maxResult(nums []int, k int) int {
	var priorityQueue maxScorePriorityQueue

	lastMaxScore := 0
	for idx, num := range nums {
		prevMax := emptyScore
		for priorityQueue.Len() > 0 {
			prevMax = priorityQueue.Peek()
			if idx <= prevMax.Index+k {
				break
			}
			priorityQueue.Poll()
		}

		lastMaxScore = prevMax.MaxScore + num
		priorityQueue.Offer(&score{Index: idx, MaxScore: lastMaxScore})
	}

	return lastMaxScore
}

type maxScorePriorityQueue struct{ scores []*score }

func (q *maxScorePriorityQueue) Len() int {
	return len(q.scores)
}

func (q *maxScorePriorityQueue) Offer(s *score) {
	offeredIdx := sort.Search(len(q.scores), func(idx int) bool {
		return q.scores[idx].MaxScore >= s.MaxScore
	})
	if offeredIdx >= len(q.scores) {
		q.scores = append(q.scores, s)
		return
	}

	q.scores = append(q.scores, nil)
	copy(q.scores[offeredIdx+1:], q.scores[offeredIdx:])
	q.scores[offeredIdx] = s
}

func (q *maxScorePriorityQueue) Peek() *score {
	if len(q.scores) == 0 {
		return emptyScore
	}
	return q.scores[len(q.scores)-1]
}

func (q *maxScorePriorityQueue) Poll() *score {
	if len(q.scores) == 0 {
		return emptyScore
	}
	old := q.scores[len(q.scores)-1]
	q.scores = q.scores[:len(q.scores)-1]
	return old
}

type score struct{ Index, MaxScore int }

var emptyScore = &score{Index: -1}
