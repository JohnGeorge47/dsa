package dsalgo

//https://leetcode.com/problems/house-robber/description/

/*
So first we need to figure out the recursive solution basically a formula
to make this work we need to keep these things in mind while figuring out the solution

So at any house i we can either rob that house or choose not to rob and instead rob the next one
sort of you can rob i-2,i,i+2 or i-1,i+1
Now we need to find the maximum at a point so at i we can choose Max(i+(i-2),i-1)
now lets try dp since we need to find a sum and we know we can memotize the solution
*/

func rob(nums []int) int {
	//Dp is hard ngl again remember the recurrence relation
	n2 := 0 // 2 house behind
	n1 := 0 // 1 house behind
	for _, num := range nums {
		maxAtn := max(num+n2, n1)
		n2 = n1
		n1 = maxAtn
	}
	return n1
}

func recursiveRob(i int, nums []int) int {
	if i < 0 {
		return 0
	}
	return max(recursiveRob(i-2, nums)+nums[i], recursiveRob(i-1, nums))
}
