package dsalgo

import (
	"math"
)

func minEatingSpeed(piles []int, h int) int {
	maxVal := piles[0]
	for _, pile := range piles {
		if pile > maxVal {
			maxVal = pile
		}
	}
	minimumSpeed := math.MaxInt
	low := 1
	high := maxVal
	for low < high {
		mid := (low + high) / 2
		timeToEat := TimeToEatPile(piles, mid)
		if timeToEat > h {
			low = mid + 1
		} else {
			high = mid - 1
			if minimumSpeed < mid {
				minimumSpeed = mid
			}
		}
	}
	return minimumSpeed
}

func TimeToEatPile(piles []int, k int) int {
	timeToEat := 0
	for _, pile := range piles {
		div := pile / k
		if pile%k != 0 {
			div = div + 1
		}
		timeToEat = timeToEat + div
	}
	return timeToEat
}
