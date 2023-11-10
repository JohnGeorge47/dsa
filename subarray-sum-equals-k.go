package dsalgo

func subarraySum(nums []int, k int) int {
	sum := 0
	count := 0
	indexSum := make(map[int]int)
	indexSum[0] = 1
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if _, ok := indexSum[sum-k]; ok {
			count = count + indexSum[sum-k]
		}
		indexSum[sum] = indexSum[sum] + 1
	}
	return count
}
