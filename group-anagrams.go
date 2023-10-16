package dsalgo

import "sort"

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	angramMap := make(map[string][]string)
	for _, str := range strs {
		charArray := []rune(str)
		sort.SliceIsSorted(charArray, func(i, j int) bool {
			return charArray[i] < charArray[j]
		})
		angramMap[string(charArray)] = append(angramMap[string(charArray)], str)
	}
	for _, strings := range angramMap {
		result = append(result, strings)
	}
	return result
}
