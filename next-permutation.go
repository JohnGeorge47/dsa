package dsalgo

//https://leetcode.com/problems/next-permutation/description/

/*
This is a very weird mathy problem
https://www.nayuki.io/page/next-lexicographical-permutation-algorithm
First you need to go from the right side of the array and find the non increasing
element
Now again go from the right side and you will need to swap the element you got with
the next greatest element
And now after swapping from the element to the end you will reverse the array
*/
func nextPermutation(nums []int) {
	pivot := findPivot(nums) - 1
	incrIndex := swapPivotElement(nums, pivot)
	if incrIndex != -1 {
		nums[incrIndex], nums[pivot] = nums[pivot], nums[incrIndex]
	}
	nums = reverseFromPivot(nums, pivot+1)
}

func findPivot(nums []int) int {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			return i
		}
	}
	return 0
}

func swapPivotElement(nums []int, pivot int) int {
	for i := len(nums); i > 0; i-- {
		if nums[i] > nums[pivot] {
			return i
		}
	}
	return -1
}

func reverseFromPivot(nums []int, pivot int) []int {
	end := len(nums) - 1
	for pivot < end {
		nums[pivot], nums[end] = nums[end], nums[pivot]
		pivot++
		end--
	}
	return nums
}
