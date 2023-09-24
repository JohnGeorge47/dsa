package dsalgo

type Trie struct {
	EndsWith bool
	WordMap  map[string]*Trie
}

func Constructor() Trie {
	return Trie{
		EndsWith: false,
		WordMap:  make(map[string]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	for _, w := range word {
		if _, ok := this.WordMap[string(w)]; !ok {
			this.WordMap[string(w)] = &Trie{
				EndsWith: false,
				WordMap:  make(map[string]*Trie),
			}
		}
		this = this.WordMap[string(w)]
	}
	this.EndsWith = true
}

func (this *Trie) Search(word string) bool {
	for _, w := range word {
		if _, ok := this.WordMap[string(w)]; !ok {
			return false
		}
		this = this.WordMap[string(w)]
	}
	return this.EndsWith
}

func (this *Trie) StartsWith(prefix string) bool {
	for _, w := range prefix {
		if _, ok := this.WordMap[string(w)]; !ok {
			return false
		}
		this = this.WordMap[string(w)]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
