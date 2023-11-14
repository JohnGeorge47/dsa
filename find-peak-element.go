package dsalgo

//https://leetcode.com/problems/find-peak-element/description/
//Binary search variation

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := low + (low+high)/2
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid] < nums[mid+1] {
			low = mid + 1
		} else if nums[mid] < nums[mid-1] {
			high = mid - 1
		}
	}
	return -1
}
