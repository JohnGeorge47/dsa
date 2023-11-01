package dsalgo

import "fmt"

// https://leetcode.com/problems/combination-sum/
/*
eg: [2,3,6,7]
Think of this like a tree alright the issue with just normally recursively traversing
this array would be you will have duplicate results for eg [2,2,3] and [3,2,2] no we need to
remove duplicates we do this by making a decision tree in this way (we choose or not choose an element)
On one side we choose an element on other decision we dont choose that element n move forward
                   []
                /      \
              [2 ]        []
             /  \         /    \
         [2,2]   [2,3]    [7]   []
*/

func combinationSum(candidates []int, target int) [][]int {
	//We create an array of the current element
	var result [][]int
	currArray := make([]int, 0)
	dfsbacktrackComb(0, 0, target, currArray, candidates, &result)
	return result
}

/*
Memory limit exceeded
*/
func dfs(i, total, target int, currArray, candidates []int, result [][]int) {
	//If the total is correct then we append to array result and return
	if total == target {
		result = append(result, currArray)
		return
	}
	// Once the out of bound conditions happen we will return
	if i > len(candidates)-1 {
		return
	}
	if total > target {
		return
	}
	currArray = append(currArray, candidates[i])
	dfs(i+1, total+candidates[i], target, currArray, candidates, result)
	dfs(i, total, target, currArray[:len(currArray)-1], candidates, result)
}

/*
We need to sort of do a better solution than continious recuriosn which exceeded mem
*/
func dfsbacktrackComb(idx, total, target int, currArray, candidates []int, result *[][]int) {
	fmt.Println(currArray)
	if total == target {
		fmt.Println(target)
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
		dfsbacktrackComb(i, total+candidates[i], target, append(append(make([]int, 0), currArray...), candidates[i]), candidates, result)
	}
}
