package dsalgo

//https://leetcode.com/problems/minimum-window-substring/

/*
This is a little more confusing than the other substring problems i have solved before
We as usual would need to maintain a map of the smaller string t
s="ADOBECODEBANC" t="ABC"
We know this is a sliding window problem and we will evict elements but what will be the condition
we will use for eviction
after we create the tmap we should get a list of all unique elements in t because our window
should have atleast these elements present
We should also keep a track of number of number of unique elements belonging to t we have in our window
*/

func minWindow(s string, t string) string {
	tMap := make(map[byte]int)
	wMap := make(map[byte]int)
	for _, i2 := range []byte(t) {
		tMap[i2] = tMap[i2] + 1
	}
	have := len(tMap)
	need := 0
	resLIndex := 0
	resRindex := 0
	//since minimum we will track max int length
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	resLen := MaxInt
	l := 0
	//We will move from beginning to end and the variable l will be used to
	//track the length from left to current position
	for r, _ := range []byte(s) {
		//We add every element we see to our map for the window
		wMap[s[r]] = wMap[s[r]] + 1
		_, ok := wMap[s[r]]
		//If the element exists in our tmap and the counts of both maps
		//For that element is the same we will do an equality check and if truue
		//have count is increased
		if ok && wMap[s[r]] == tMap[s[r]] {
			have = have + 1
		}
		//This is condition where our current window has all the needed elements in t
		if have == need {
			//To find the minimum substring
			if r-l+1 < resLen {
				resRindex = r
				resLIndex = l
				resLen = r - l + 1
			}
			//We start evicting elements from the left of the window
			wMap[s[l]] = wMap[s[l]] - 1
			_, ok := tMap[s[l]]
			if ok && wMap[s[l]] < tMap[s[r]] {
				have = have - 1
			}
			l = l + 1
		}
	}
	return s[resLIndex:resRindex]
}
