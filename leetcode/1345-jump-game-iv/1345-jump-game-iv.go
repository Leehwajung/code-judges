package main

func minJumps(arr []int) int {
	numIdxs := make(map[int][]int)
	for idx := len(arr) - 1; idx >= 0; idx-- {
		num := arr[idx]
		numIdxs[num] = append(numIdxs[num], idx)
	}

	queue := []int{0}
	visited := make(map[int]bool)
	jumps := -1

	for len(queue) > 0 {
		jumps++
		loops := len(queue)

		for i := 0; i < loops; i++ {
			currIdx := queue[0]
			queue = queue[1:]
			if currIdx == len(arr)-1 {
				return jumps
			}

			if currIdx < 0 || currIdx >= len(arr) || visited[currIdx] {
				continue
			}

			visited[currIdx] = true

			nextIdx := currIdx + 1
			if nextIdx < len(arr) && arr[nextIdx] != arr[currIdx] {
				queue = append(queue, nextIdx)
			}

			prevIdx := currIdx - 1
			if prevIdx >= 0 && arr[prevIdx] != arr[currIdx] {
				queue = append(queue, prevIdx)
			}

			queue = append(queue, numIdxs[arr[currIdx]]...)
			delete(numIdxs, arr[currIdx])
		}
	}

	return jumps
}
