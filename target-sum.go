package dsalgo

func findTargetSumWays(nums []int, target int) int {
	count := 0
	CalculateSum(0, target, nums, &count, 0)
	return count
}

func CalculateSum(currSum int, target int, nums []int, count *int, idx int) {
	if idx > len(nums)-1 {
		return
	}
	if currSum == target {
		*count = *count + 1
	}
	CalculateSum(currSum+nums[idx], target, nums, count, idx+1)
	CalculateSum(currSum-nums[idx], target, nums, count, idx+1)
}
