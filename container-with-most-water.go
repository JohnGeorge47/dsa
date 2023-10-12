package dsalgo

func maxArea(height []int) int {
	max := 0
	left := 0
	right := len(height) - 1
	for left < right {
		area := (right - left) * minima(height[left], height[right])
		if area > max {
			max = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

func minima(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxima(a, b int) int {
	if a > b {
		return a
	}
	return b
}
