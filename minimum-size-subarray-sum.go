package dsalgo

import "math"

func minSubArrayLen(target int, nums []int) int {
	minSize := math.MaxInt
	numsStart := 0
	sum := nums[0]
	if sum == target {
		minSize = min(minSize, 1)
	}
	for i := 1; i < len(nums); i++ {
		sum = sum + nums[i]
		if sum == target {
			minSize = min(minSize, i-numsStart+1)
		}
		if sum > target {
			for sum > target {
				sum = sum - nums[numsStart]
				if sum == target {
					minSize = min(minSize, i-numsStart+1)
				}
				numsStart = numsStart + 1
			}
		}
	}
	return minSize
}
