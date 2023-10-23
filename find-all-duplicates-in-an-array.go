package dsalgo

/*
Read the problem there is at max two occurances and
the values are between 1<i<n
this means if we mark a position as -ve that means if there is a duplicate we will rewrite that
position so only negative values wont have duplicates
*/
func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	for _, num := range nums {
		if nums[abs(num)-1] < 0 {
			res = append(res, abs(num))
		} else {
			nums[abs(num)-1] = -1 * nums[abs(num)-1]
		}
	}
	return res
}
