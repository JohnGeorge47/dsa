package dsalgo

//Problem:https://leetcode.com/problems/implement-trie-prefix-tree/description/

/*
What is a trie even Its i guess what you would call a prefix tree It basically
stores letters in a word in a manner that lookup is effecient. sure we can use a
map but it wont allow for prefix checking
lets say we insert "cat" our tree will look like
Note:Eow is end of word
			c
            |
            a
            |
            t eow->true
now lets say we insert "car"
             c
           /
          a
         / \
      t(eow)     r (eow)
now we insert "dog"
              c    d
           /        \
          a          o
         / \          \
      t(eow) r(eow)    g(eow)
*/

// The trie has a map of words at that node as well as a boolean signifying if
// that node will be the end of a word
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

/*
While inserting we go throught the entire word
*/
func (this *Trie) Insert(word string) {
	//We will go letter by letter we will check in that nodes map if the letter
	//exists and if not create a new node which will be refering to that letter
	//in map if not we will continue to next node which is mapped to that letter
	for _, w := range word {
		if _, ok := this.WordMap[string(w)]; !ok {
			this.WordMap[string(w)] = &Trie{
				EndsWith: false,
				WordMap:  make(map[string]*Trie),
			}
		}
		this = this.WordMap[string(w)]
	}
	//Once we reach the end we mark the node as the end
	this.EndsWith = true
}

func (this *Trie) Search(word string) bool {
	for _, w := range word {
		//if word does not exist at the particular level we return false
		if _, ok := this.WordMap[string(w)]; !ok {
			return false
		}
		this = this.WordMap[string(w)]
	}
	//if we reach the end then we return based on value
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
