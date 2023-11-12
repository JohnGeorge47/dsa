package dsalgo

func longestOnes(nums []int, k int) int {
	ones := 0
	maxLen := 0
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			ones = ones + 1
		}
		currLength := i - left - ones + 1
		if currLength > k && left < len(nums) {
			if nums[left] == 1 {
				ones = ones - 1
			}
			left = left + 1
			currLength = i - left - ones + 1
		}
		maxLen = max(maxLen, i-left+1)
	}

	return maxLen
}
