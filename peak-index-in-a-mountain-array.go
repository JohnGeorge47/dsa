package dsalgo

func peakIndexInMountainArray(arr []int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] < arr[mid+1] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
