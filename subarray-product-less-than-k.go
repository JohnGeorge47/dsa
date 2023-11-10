package dsalgo

func numSubarrayProductLessThanK(nums []int, k int) int {
	slow := 0
	result := 0

	windowMult := 1
	for i := 0; i < len(nums); i++ {
		windowMult = windowMult * nums[i]
		for slow <= i && windowMult > k {
			windowMult = windowMult / nums[slow]
			slow++
		}
		result = result + i - slow + 1
	}
	return result
}
