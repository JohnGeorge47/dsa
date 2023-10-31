package dsalgo

type Position struct {
	prevVal int
	idx     int
}

//https://leetcode.com/problems/longest-increasing-subsequence/description/
/*
Didnt work turned out to be a dp problem
1.First lets look at a brute force approach of 2^n

		We look at each position then we add the element if the element
		if its greater than current element and like this we find all substrings
		and check greatest one
	 2. Next lets look at memotization we see that some paths are repeated in the decision
	    tree we add this to a hash map, what we notice is the last last element does not have
	    anything after and will always be 1 we can use this to do a dp solution
	    and then
*/
func lengthOfLIS1(nums []int) int {
	maxLen := 0
	p := Position{
		prevVal: nums[len(nums)-1],
		idx:     len(nums) - 1,
	}
	currLen := 0
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < p.prevVal {
			p.prevVal = nums[i]
			currLen = p.idx - i
		} else {
			p.prevVal = nums[i]
			p.idx = i
			currLen = 1
		}
		if currLen > maxLen {
			maxLen = currLen
		}
	}
	return maxLen
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[len(nums)-1] = 1
	res := 0
	if len(nums) == 1 {
		return 1
	}
	for i := len(dp) - 2; i >= 0; i-- {
		if dp[i] == 0 {
			dp[i] = 1
		}
		for j := i + 1; j < len(dp); j++ {
			if nums[i] < nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
