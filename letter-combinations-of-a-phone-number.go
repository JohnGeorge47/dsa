package dsalgo

/*
This is another permutation combination type problem so again its backtracking
so you will need to do combination of each digit and all its possible values then recursively call with next digit
*/
func letterCombinations(digits string) []string {
	allCombos := make([]string, 0)
	if digits == "" {
		return []string{""}
	}
	m := map[byte][]string{
		'0': []string{"0"},
		'1': []string{"1"},
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
	recursiveCombo(digits, 0, "", &allCombos, m)
	return allCombos
}

func recursiveCombo(digits string, idx int, combo string, allCombo *[]string, m map[byte][]string) {
	if len(combo) == len(digits) {
		*allCombo = append(*allCombo, combo)
		return
	}
	for _, s := range m[digits[idx]] {
		recursiveCombo(digits, idx+1, combo+s, allCombo, m)
	}
}
