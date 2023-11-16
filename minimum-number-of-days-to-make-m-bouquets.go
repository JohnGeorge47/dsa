package dsalgo

func minDays(bloomDay []int, m int, k int) int {
	maxNumber := 0
	for _, i2 := range bloomDay {
		if i2 > maxNumber {
			maxNumber = i2
		}
	}
	lowDays := 0
	for lowDays < maxNumber {
		mid := lowDays + (maxNumber-lowDays)/2

	}
}

func FindBoq(day int, bloomDay []int) bool {

}
