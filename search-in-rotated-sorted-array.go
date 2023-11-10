package dsalgo

func search(nums []int, target int) int {
	low := nums[0]
	high := nums[len(nums)-1]
	mid := (low + high) / 2
	for low <= high {
		if nums[mid] == target {
			return mid
		}
		//Left sorted search
		if nums[low] <= nums[mid] {
			if nums[low] <= target && target <= nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
