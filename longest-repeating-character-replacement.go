package dsalgo

/*
Remember you need to replace elements within a substring
the main equation you need to know is
number_to_replace=length-most_frequent
in our case
length_substr-most_frequent <=k
remember if a new element is added we can only move left pointer and change subtr length
*/
func characterReplacement(s string, k int) int {
	wordArr := make([]int, 26)
	maxFreq := 0
	result := 0
	left := 0
	for i := 0; i < len(s); i++ {
		wordArr[s[i]-'A']++
		maxFreq = max(maxFreq, wordArr[s[i]-'A'])
		charToReplace := i - left + 1 - maxFreq
		for charToReplace > k {
			wordArr[s[left]-'A']--
			left++
		}
		result = max(result, i-left+1)
	}
	return result
}
