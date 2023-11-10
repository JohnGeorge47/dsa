package dsalgo

/*
Was asked in an interview didnt understand how but it is a sliding window problem

Basically you maintain a map of k distinct characters each time you encounter
a unique char you add it to the map once your map has reached its size=k you start
removing items from left of string in the map until you again reach k and then  proceed
with the traversal
You will need two pointers for this
*/
func kdistinctCharacters(s string, k int) {
	charMap := make(map[uint8]int)
	left := 0
	right := 0
	currMax := 0
	for right < len(s) {
         charMap[s[right]]++
		}
		currMax = max(currMax, right-left+1)
	}
}
