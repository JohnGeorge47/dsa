package dsalgo

import "sort"

//https://leetcode.com/problems/3sum/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			if nums[i] == nums[i-1] {
				continue
			}
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[l] + nums[r] + nums[i]
			if sum == 0 {
				res = append(res, []int{l, r, i})
				l = l + 1
				for l-1 == l && l < r {
					l = l + 1
				}
			}
			if sum > 0 {
				r = r - 1
			}
			if sum < 0 {
				l = l + 1
			}
		}
	}
	return result
}
