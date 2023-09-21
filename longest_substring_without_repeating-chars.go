package dsalgo

func lengthOfLongestSubstring(s string) int {
	max := 0
	currMax := 0
	start := 0
	charMap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		if val, ok := charMap[string(s[i])]; ok {
			if !(val < start) {
				start = val + 1
				currMax = i - start + 1
			} else {
				currMax = i - start + 1
			}
		} else {
			currMax = i - start + 1
		}
		if currMax > max {
			max = currMax
		}
		charMap[string(s[i])] = i
	}
	return max
}
