package dsalgo

//https://leetcode.com/problems/fruit-into-baskets/

type PositionVal struct {
	LastPosition int
	Fruit        int
}

/*
Failed for 1 testcase
*/
func totalFruit2(fruits []int) int {
	basketMap := make(map[int]int)
	result := 1
	currMax := 1
	basketMap[fruits[0]] = 1
	firstSeen := fruits[0]

	for i := 1; i < len(fruits); i++ {
		if len(basketMap) == 2 {
			if _, ok := basketMap[fruits[i]]; !ok {
				currMax = currMax - basketMap[firstSeen]
				delete(basketMap, firstSeen)
				for k, _ := range basketMap {
					firstSeen = k
				}
				if firstSeen == fruits[i] {
					for k, _ := range basketMap {
						if k != firstSeen {
							firstSeen = k
						}
					}
				}
				basketMap[fruits[i]] = 1
				currMax = currMax + 1
			} else {
				basketMap[fruits[i]] = basketMap[fruits[i]] + 1
				currMax = currMax + 1
			}
		} else {
			basketMap[fruits[i]] = basketMap[fruits[i]] + 1
			currMax = currMax + 1
		}
		result = max(currMax, result)
	}
	return result
}

func totalFruit1(fruits []int) int {
	basketMap := make(map[int]int)
	result := 1
	startCount := 0
	for i := 0; i < len(fruits); i++ {
		basketMap[fruits[i]] = basketMap[fruits[i]] + 1
		for len(basketMap) > 2 {
			basketMap[fruits[startCount]] = basketMap[fruits[startCount]] - 1
			if basketMap[startCount] == 0 {
				delete(basketMap, basketMap[startCount])
				startCount = startCount + 1
			}
			startCount = startCount + 1
		}
		result = max(result, i-startCount+1)
	}
	return result
}
