package containerwithmostwater

func maxArea(height []int) int {
	l := 0
	h := len(height) - 1

	var maxArea int
	var area int

	for l < h {
		area = min(height[l], height[h]) * abs(l-h)
		if maxArea < area {
			maxArea = area
		}

		if height[l] < height[h] {
			l++
		} else {
			h--
		}

	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
