package dsalgo

//https://leetcode.com/problems/combination-sum-iii/submissions/

func combinationSum3(k int, n int) [][]int {
	inpArray := make([]int, 9)
	var result [][]int
	for i := 1; i < 10; i++ {
		inpArray[i] = i
	}
	dfcomb3(k, 0, n, 0, inpArray, make([]int, 0), &result)
	return result
}

func dfcomb3(arrSize, currSum, target, idx int, inpArr, currArr []int, result *[][]int) {
	if currSum == target && len(currArr) == arrSize {
		*result = append(*result, currArr)
	}
	if currSum > target {
		return
	}
	if len(currArr) > arrSize {
		return
	}
	for i := idx; i < len(inpArr); i++ {
		dfcomb3(arrSize, currSum+inpArr[i], target, i+1, inpArr, append(append(make([]int, 0), currArr...), inpArr[i]), result)
	}
}
