package dsalgo

type TimeMap struct {
	ValueMap map[string]*KeyNode
}

type KeyNode struct {
	keyMap  map[int]string
	timeVal []int
}

//func Constructor() TimeMap {
//	return TimeMap{ValueMap: make(map[string]*KeyNode)}
//}

func findClosest(arr []int, target int) int {
	l := 0
	r := len(arr) - 1
	for l < r {
		if abs(arr[l]-target) <= abs(arr[r]-target) {
			r--
		} else {
			l--
		}
	}
	return l
}

func abs(val int) int {
	if val < 0 {
		return -1 * val
	}
	return val
}

func (this *TimeMap) Set(key string, value string, timestamp int) {

}

func (this *TimeMap) Get(key string, timestamp int) string {

}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
