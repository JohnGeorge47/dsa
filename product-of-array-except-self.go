package dsalgo

//Problem:https://leetcode.com/problems/product-of-array-except-self/description/
/*
I found this problem really hard to understand
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Now theres one solution which says imagine the product of 2 arrays which is
the prefix array (one being product of all elements till that element) and
postfix(reverse of this which is product of all elements after that array) now if
you multiply all the elements in the same position in these two arrays you will get
product except the array itself. I saw this in a yt comment which made it much easier
visualize

lets replace array with letters =[a,b,c,d]

prefix:[  a  ,  a*b   , a*b*c ,a*b*c*d ]
postfix: [a*b*c*d, b*c*d, c*d, d]

now multiply the two excluding the element and you will get your answer
*/

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	for i, _ := range result {
		result[i] = 1
	}
	curr := nums[0]
	for i := 1; i < len(nums); i++ {
		result[i] = curr
		curr = nums[i] * curr
	}
	prev := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * prev
		prev = nums[i] * prev
	}
	return result
}
