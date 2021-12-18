package main

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	} else if len(height) == 2 {
		return intMin(height[0], height[1])
	}

	max, maxHeight, maxLeftHeight, maxRightHeight := 0, 0, 0, 0
	for left, leftHeight := range height[:len(height)-1] {
		if leftHeight < maxLeftHeight {
			continue
		}
		maxLeftHeight = leftHeight
		maxRightHeight = 0

		for right := len(height) - 1; right > left; right-- {
			if height[right] < maxRightHeight {
				continue
			}
			maxRightHeight = height[right]

			currHeight := intMin(leftHeight, height[right])
			if currHeight < maxHeight {
				continue
			}

			area := currHeight * (right - left)
			if max < area {
				max = area
			}
		}
	}
	return max
}

func intMin(x, y int) int {
	if x > y {
		return y
	}
	return x
}
