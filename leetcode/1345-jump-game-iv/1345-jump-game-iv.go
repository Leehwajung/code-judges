package main

func minJumps(arr []int) int {
	numIdxs := make(map[int][]int)
	for idx := len(arr) - 1; idx >= 0; idx-- {
		num := arr[idx]
		numIdxs[num] = append(numIdxs[num], idx)
	}

	visited := make(map[int]bool)

	type idxJumps struct{ idx, jumps int }
	queue := []*idxJumps{{idx: 0, jumps: 0}}

	for ; len(queue) > 0; queue = queue[1:] {
		curr := queue[0]
		if curr.idx == len(arr)-1 {
			return curr.jumps
		}

		if curr.idx < 0 || curr.idx >= len(arr) || visited[curr.idx] {
			continue
		}

		visited[curr.idx] = true
		nextJumps := curr.jumps + 1

		nextIdx := curr.idx + 1
		if nextIdx < len(arr) && arr[nextIdx] != arr[curr.idx] {
			queue = append(queue, &idxJumps{idx: nextIdx, jumps: nextJumps})
		}

		prevIdx := curr.idx - 1
		if prevIdx >= 0 && arr[prevIdx] != arr[curr.idx] {
			queue = append(queue, &idxJumps{idx: prevIdx, jumps: nextJumps})
		}

		for _, jumpIdx := range numIdxs[arr[curr.idx]] {
			queue = append(queue, &idxJumps{idx: jumpIdx, jumps: nextJumps})
		}
		delete(numIdxs, arr[curr.idx])
	}

	return len(arr) - 1
}
