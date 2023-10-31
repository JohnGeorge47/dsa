package dsalgo

type CacheKey struct {
	idx int
	str string
}

/*
Definately not the best solution but its accepted
*/
func countSubstrings(s string) int {
	cache := make(map[CacheKey]struct{})
	for i, _ := range s {
		l := i
		r := i
		//odd check
		for l <= r && r < len(s) && l >= 0 && s[l] == s[r] {
			cache[CacheKey{
				idx: i,
				str: s[l : r+1],
			}] = struct{}{}
			l--
			r++
		}
		l = i
		r = i + 1
		for l <= r && r < len(s) && l >= 0 && s[l] == s[r] {
			cache[CacheKey{
				idx: i,
				str: s[l : r+1],
			}] = struct{}{}
			l--
			r++
		}
	}
	return len(cache)
}
