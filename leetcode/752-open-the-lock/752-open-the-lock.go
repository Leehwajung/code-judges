package main

const wheels = 4

func openLock(deadends []string, target string) int {
	targetVal := parseLockValue(target)
	deadendsSet := make(map[[wheels]uint8]struct{})
	for _, deadend := range deadends {
		deadendsSet[parseLockValue(deadend)] = struct{}{}
	}

	type queueItem = struct {
		val  [wheels]uint8
		turn int
	}
	queue := []queueItem{{
		val:  [wheels]uint8{0, 0, 0, 0},
		turn: 0,
	}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.val == targetVal {
			return curr.turn
		} else if _, ok := deadendsSet[curr.val]; ok {
			continue
		} else {
			deadendsSet[curr.val] = struct{}{}
		}

		for wheel := uint8(0); wheel < wheels; wheel++ {
			{
				rotatedVal := curr.val
				rotateLockValueForward(&rotatedVal, wheel)
				queue = append(queue, queueItem{
					val:  rotatedVal,
					turn: curr.turn + 1,
				})
			}
			{
				rotatedVal := curr.val
				rotateLockValueBackward(&rotatedVal, wheel)
				queue = append(queue, queueItem{
					val:  rotatedVal,
					turn: curr.turn + 1,
				})
			}
		}
	}

	return -1
}

func parseLockValue(s string) (result [wheels]uint8) {
	for i, digit := range s {
		result[i] = uint8(digit - '0')
	}
	return
}

func rotateLockValueForward(val *[wheels]uint8, wheel uint8) {
	if val[wheel] == 9 {
		val[wheel] = 0
	} else {
		val[wheel]++
	}
}

func rotateLockValueBackward(val *[wheels]uint8, wheel uint8) {
	if val[wheel] == 0 {
		val[wheel] = 9
	} else {
		val[wheel]--
	}
}
