package dsalgo

//https://leetcode.com/problems/palindrome-partitioning/

/*
There are few things to consider here
Once we start looking at a string we need to check if it is a palindrome
if not we can abort all other suffix string combos
so basically we can split the string to a prefix and a suffix prefix being substring till that
character
suffix is from character to end of string so with each iteration we just reduce substr size by 1
*/
func partition(s string) [][]string {
	var result [][]string
	palindromeDfs(s, make([]string, 0), &result)
	return result
}

func palindromeDfs(s string, currStrArr []string, result *[][]string) {
	if s == "" || len(s) == 0 {
		*result = append(*result, currStrArr)
	}
	for i := 1; i <= len(s); i++ {
		if !isPalindrome(s[:i]) {
			continue
		}
		palindromeDfs(s[i:], append(append(make([]string, 0), currStrArr...), s[:i]), result)
	}
	return
}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	mid := len(s) / 2
	if len(s)%2 != 0 {
		l := mid
		r := mid
		for l >= 0 && r < len(s) {
			if s[l] != s[r] {
				return false
			}
			l--
			r++
		}
		return true
	}
	l := mid - 1
	r := mid
	for l >= 0 && r < len(s) {
		if s[l] != s[r] {
			return false
		}
		l--
		r++
	}
	return true
}
