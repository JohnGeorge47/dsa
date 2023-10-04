package dsalgo

//https://leetcode.com/problems/partition-equal-subset-sum/

/*
Pretty much dp but the brute force is add every element and see if it equals the
value we need which is half of sum
for eg [1,5,11,5] the sum is 22 but equally partitioned means there should be 2
so each partition will have sum of 11
so we can sum each combination using dfs till sum is target or array size
is len(arr)
or do a sort of dp with a set where we keep the sum of each element
*/
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum = sum + num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	sumSet := make(map[int]struct{})
	sumSet[0] = struct{}{}
	for _, num := range nums {
		for i, _ := range sumSet {
			if i+num == target {
				return true
			}
			sumSet[i+num] = struct{}{}
		}
	}
	return false
}
