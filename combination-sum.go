package dsalgo

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

var result [][]int

func combinationSum(candidates []int, target int) [][]int {
	//We create an array of the current element
	currArray := make([]int, 0)
	dfs(0, 0, target, currArray, candidates)
	return result
}

func dfs(i, total, target int, currArray, candidates []int) {
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
	dfs(i+1, total+candidates[i], target, currArray, candidates)
	dfs(i, total, target, currArray[:len(currArray)-1], candidates)
}
