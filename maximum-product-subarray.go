package dsalgo

//https://leetcode.com/problems/maximum-product-subarray/

/*
DP wow who wouldve thought

So we need to understand three things

1. If the elements are all +ve  we just return entire array
2. If there are -ve elements.
   a. Even negative means we return entire product
   b. Odd means we exclude that number
3. When value is 0 we start a new subarray

In the below code we maintain a max and a min
This is due to occurence of -ve numbers
max till a point is num[i]*max
but if num[i] is -ve?
then max at the point is min

We use min to track the -ve values and flip in case we see another negative
*/

func maxProduct(nums []int) int {
	result := nums[0]
	maximum := nums[0]
	minimum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maximum, minimum = minimum, maximum
		}
		maximum = maxima(maximum*nums[i], nums[i])
		minimum = min(minimum*nums[i], nums[i])
		result = maxima(result, maximum)
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxima(a, b int) int {
	if a > b {
		return a
	}
	return b
}
