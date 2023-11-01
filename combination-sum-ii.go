package dsalgo

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	//We create an array of the current element
	var result [][]int
	currArray := make([]int, 0)
	sort.Ints(candidates)
	dfsbacktrackComb2(0, 0, target, currArray, candidates, &result)
	return result
}

func dfsbacktrackComb2(idx, total, target int, currArray, candidates []int, result *[][]int) {
	if total == target {
		*result = append(*result, currArray)
		return
	}
	if idx > len(candidates) {
		return
	}
	if total > target {
		return
	}
	for i := idx; i < len(candidates); i++ {
		if i > idx && candidates[i] == candidates[i-1] {
			continue
		}
		dfsbacktrackComb2(i+1, total+candidates[i], target, append(append(make([]int, 0), currArray...), candidates[i]), candidates, result)
	}
}
