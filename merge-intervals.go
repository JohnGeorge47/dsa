package dsalgo

func merge(intervals [][]int) [][]int {
	var result [][]int
	currInterval := intervals[0]
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		if currInterval[1] >= interval[0] && interval[1] >= currInterval[0] {
			currInterval = []int{min(currInterval[0], interval[0]), max(interval[1], currInterval[1])}
		} else {
			result = append(result, currInterval)
			currInterval = interval
		}
	}
	result = append(result, currInterval)
	return result
}
