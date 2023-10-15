package dsalgo

//https://leetcode.com/problems/permutation-in-string/description/

/*
I think im getting a hang of this sliding window thing keep in mind this especially
We can do this two ways one is by maintaining an array or a map
An array works because we have only 26 letters so an array of len 26 with each position
having count of each character in string
We use sliding window of size of smaller string and we add all elements in array uptill len(s1)
in the windowMap or windowArray in our case
Then we start evicting elements as we move sliding window along and add newest element
then we check for equality between the elements in the two arrays
*/
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	occArr := make([]int, 26)
	for _, i2 := range s1 {
		occArr[i2-'a'] = occArr[i2-'a'] + 1
	}
	windowSize := len(s1)
	windArr := make([]int, 26)
	for _, i2 := range s2[:windowSize] {
		windArr[i2-'a'] = windArr[i2-'a'] + 1
	}
	if isEqual(occArr, windArr) {
		return true
	}
	for i := windowSize; i < len(s2); i++ {
		windArr[s2[i-windowSize]-'a'] = windArr[s2[i-windowSize]-'a'] - 1
		windArr[s2[i]-'a'] = windArr[s2[i]-'a'] + 1
		if isEqual(windArr, occArr) {
			return true
		}
	}
	return false
}

func isEqual(pDict, sDict []int) bool {
	for i := 0; i < 26; i++ {
		if pDict[i] != sDict[i] {
			return false
		}
	}
	return true
}
