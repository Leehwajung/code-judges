package main

func canReach(arr []int, start int) bool {
	visited := make(map[int]bool)
	return canReachRecursive(arr, start, visited)
}

func canReachRecursive(arr []int, idx int, visited map[int]bool) bool {
	if idx < 0 || idx >= len(arr) || visited[idx] {
		return false
	}

	visited[idx] = true

	if arr[idx] == 0 {
		return true
	}

	if canReachRecursive(arr, idx-arr[idx], visited) {
		return true
	}
	if canReachRecursive(arr, idx+arr[idx], visited) {
		return true
	}
	return false
}
