package dsalgo

/*
https://leetcode.com/problems/generate-parentheses/description/
*/

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	return result
}

func backTrackParentesis(n, left, right int, currStr string, result *[]string) {
	if len(currStr) == 2*n {
		*result = append(*result, currStr)
		return
	}
	if left < n {
		backTrackParentesis(n, left+1, right, currStr+"(", result)
	}
	if right < left {
		backTrackParentesis(n, left, right+1, currStr+")", result)
	}

}
