package dsalgo

func maxVowels(s string, k int) int {
	vowels := make(map[byte]struct{})
	vowels['a'] = struct{}{}
	vowels['i'] = struct{}{}
	vowels['e'] = struct{}{}
	vowels['o'] = struct{}{}
	vowels['u'] = struct{}{}
	currSize := 0
	for i := 0; i < k; i++ {
		if _, ok := vowels[s[i]]; ok {
			currSize = currSize + 1
		}
	}
	maxSize := currSize
	for i := k; i < len(s); i++ {
		if _, ok := vowels[s[i-k]]; ok {
			currSize = currSize - 1
		}
		if _, ok := vowels[s[i]]; ok {
			currSize = currSize + 1
		}
		maxSize = max(currSize, maxSize)
	}

	return maxSize
}
