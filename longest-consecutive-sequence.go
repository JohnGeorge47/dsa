package dsalgo

//https://leetcode.com/problems/longest-consecutive-sequence/description/

/*
Think of the numberline and you need to check if element has something to left if not its not
its not the start of a subsequence
if yes you can keep iterating from the element till you dont find anything to right
*/
func longestConsecutive(nums []int) int {
	//We first maintain a map of all elements in nums
	longest := 0
	numberline := make(map[int]struct{})
	for _, num := range nums {
		numberline[num] = struct{}{}
	}
	for _, num := range nums {
		if _, ok := numberline[num-1]; !ok {
			length := 1
			for {
				if _, ok := numberline[num+length]; ok {
					length = length + 1
					continue
				}
				longest = max(length, longest)
				break
			}
		}
	}
	return longest
}
