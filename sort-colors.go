package dsalgo

//https://leetcode.com/problems/sort-colors/description/

/*
eg [2,0,2,1,1,0]

Theres a really straightforward solution called bucket sort which basically is a
2 pass solution which we got through entire array once and add count in a map and replace
elements accordingly but lets try a one pass solution its apparently
called the dutch flag partitioning or something
lets first decide that we have 2 ptrs a left and a right
and as usual run a while
but we have another pointer which is to traverse the array
since we want to sort this
we first make sure any 2 we encounter will be exchanged with last element and right--
and
any 0 we encounter will be exchanged with left but we will not increase i to manage an edge case
for eg [0 , 1, 2, 1 ,0,2]
Since we dont want stray 0's and the fact that 2's will always be at end and 1's
in middle we need to make sure that whenever a right shift we do not increase the
i pointer (one which traverses array)

	L     i     R
*/
func sortColors(nums []int) {
	left, i := 0, 0
	right := len(nums) - 1
	for i <= right {
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
			i++
		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		} else {
			i++
		}
	}
}
