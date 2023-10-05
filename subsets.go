package dsalgo

import "sort"

//https://leetcode.com/problems/subsets/

/*
Remember permutations this is like easy if you get that

this is again a backtracking problem now unlike permutations where we have the
permutation being a combination equalling length of array here
we have initially an empty (curr[]) we traverse the nums () and we add first element to curr[1]
We keep appending each element to this curr array
then do a recursive call but with i+1 so that we only traverse the next elements needed
we keep doing this for each element
*/
func subsets(nums []int) [][]int {
	subsetResult := make([][]int, 0)
	sort.Ints(nums)
	subsetBackTrack(0, make([]int, 0), nums, &subsetResult)
	return subsetResult
}

func subsetBackTrack(idx int, curr []int, nums []int, result *[][]int) {
	*result = append(*result, curr)
	for i := idx; i < len(nums); i++ {
		subsetBackTrack(
			i+1,
			append(curr, nums[i]),
			nums[i:len(nums)],
			result,
		)
	}
}
