package dsalgo

func shipWithinDays(weights []int, days int) int {
	maxCapacity := 0
	sum := 0
	for _, weight := range weights {
		if weight > maxCapacity {
			maxCapacity = weight
		}
		sum = sum + weight
	}
	left := maxCapacity
	right := sum
	for left < right {
		mid := left + (left+right)/2
		dayNeeded := DaysToShip(weights, mid)
		if dayNeeded < days {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func DaysToShip(weights []int, capacity int) int {
	days := 0
	totalWeight := 0
	for _, weight := range weights {
		totalWeight = weight + totalWeight
		if totalWeight > capacity {
			days = days + 1
			totalWeight = weight
		}
	}
	return days
}
