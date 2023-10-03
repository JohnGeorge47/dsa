package dsalgo

//https://leetcode.com/problems/permutations/

/*
 This is little confusing but its sort of like backtracking lets look at an example
 [1,2,3] so what we need to do here is we will need to realize that permutation is
 mathematically something like this _ _ _ where we can add any of the elements in these
 blank spaces so using that lets look at a graph
                []
               /
             [1]
           /     \
          [3,2]   [2,3]
*/

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	//Its important to note we need a pointer to the result since we are modifying
	//the same object
	backTrack(nums, make([]int, 0), &result)
	return result
}

// Currperm holds the permutation object which is the object too which each element is added
func backTrack(nums, currperm []int, result *[][]int) {
	//At our leaf node once we have found permutation the nums will become of 0 length
	//since all the values in num will be part of currperm and we can add the value to result
	if len(nums) == 0 {
		*result = append(*result, currperm)
		return
	}
	/*
		We loop through nums initially lets say nums=[1,2,3] since permutation is _ _ _
		we first fix 1 in the loop by removing it from nums and adding it to currperm so [1],[2,3]
		(we make copies since in go otherwise it modifies original array) and then we call the same function
		recursively, then we go to next index and we fix [2],[1,3]. and [3],[1,2]
		Now lets look at the recursive case [1],[2,3]
		currPerm=[1,2] nums=[3]  currPerm=[1,3] nums=[2]
		currPerm=[1,2,3] nums=[] currPerm=[1,3,2] nums=[] so we reach base case and return here
	*/
	for i, num := range nums {
		backTrack(
			append(append([]int{}, nums[:i]...), nums[i+1:]...),
			append(currperm, num),
			result,
		)
	}
}
