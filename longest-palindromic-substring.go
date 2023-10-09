package dsalgo

//https://leetcode.com/problems/longest-palindromic-substring/

/*
Could only come up with a brute force solution

Ok found the solution remember a palindrome is a reflection around the mid element
so there are two cases here
 1. Odd we move from centre and each element to left and right should be equal
    a  b  a   its a palindrome when left and right are equal
 2. Even we move from pointer on one side and pointer +1 on other
*/
func longestPalindrome(s string) string {
	palinMax := ""
	for i, _ := range s {
		for i3, _ := range s {
			//if isPalindrome(s[i:i3]) {
			//	if len(s[i:i3]) > len(palinMax) {
			//		palinMax = s[i:i3]
			//	}
			//}
		}
	}
	return palinMax
}

func longestPalindrome2(s string) string {
	palinMax := ""
	for i, _ := range s {
		//Odd palindrome check
		l := i
		r := i
		for l <= r && l >= 0 && r < len(s) && s[l] == s[r] {
			if len(s[l:r+1]) > len(palinMax) {
				palinMax = s[l : r+1]
			}
			r = r + 1
			l = l - 1
		}
		//Even palindrome check
		l = i
		r = i + 1
		for l <= r && l >= 0 && r < len(s) && s[l] == s[r] {
			if len(s[l:r+1]) > len(palinMax) {
				palinMax = s[l : r+1]
			}
			r = r + 1
			l = l - 1
		}
	}
	return palinMax
}
