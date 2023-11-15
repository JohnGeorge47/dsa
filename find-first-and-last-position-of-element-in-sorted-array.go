package dsalgo

import "fmt"

func searchRange(nums []int, target int) []int {
	left := binarySearch(target, nums)
	fmt.Println(left)
	return []int{-1, -1}
}

func binarySearch(target int, nums []int) int {
	low := 0
	high := len(nums) - 1
	idx := -1
	for low < high {
		//Find low range
		mid := low + (high-low)/2
		if target == nums[mid] {
			idx = mid
			high = mid - 1
		}
		if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return idx
}
