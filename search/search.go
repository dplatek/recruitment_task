package search

func FindCloseEnoughValue(data []int, value int, margin float64) (int, int) {
	left, right := 0, len(data)-1
	var closestValue int
	var closestIndex int
	minDiff := margin + 1

	for left <= right {
		mid := left + (right-left)/2
		diff := float64(abs(data[mid] - value))

		if diff <= margin {
			if diff < minDiff {
				minDiff = diff
				closestValue = data[mid]
				closestIndex = mid
			}
			if data[mid] < value {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else if data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return closestValue, closestIndex
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
