package dsalgo

// Criss crossing remember it then check max of 0,0 and min (1,1)
// to find overlaps
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	left := 0
	right := 0
	var result [][]int
	for left < len(firstList) && right < len(secondList) {
		l1 := firstList[left]
		l2 := secondList[right]
		//Check if intervals are overlapping
		if l1[1] >= l2[0] && l2[1] >= l1[0] {
			val1 := max(l2[0], l1[0])
			val2 := min(l1[1], l2[1])
			result = append(result, []int{val1, val2})
		}
		if l2[1] > l1[1] {
			left++
		} else {
			right++
		}
	}
	return result
}
