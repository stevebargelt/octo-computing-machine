package highestproduct

import (
	"errors"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// GetHighestProduct
// O(n) time and O(1) additional space.
// Be Greedy... even as comvuluted as this looks the complexity and
//	space are much better than looping inside loops
func GetHighestProduct(arry []int) (int, error) {

	if arry == nil {
		return 0, errors.New("Array cannot be nil")
	}
	if len(arry) < 3 {
		return 0, errors.New("Array has to have at least 3 elements")
	}

	highest := max(arry[0], arry[1])
	lowest := min(arry[0], arry[1])
	highestProdOf2 := arry[0] * arry[1]
	lowestProdOf2 := arry[0] * arry[1]
	highestProdOf3 := arry[0] * arry[1] * arry[2]
	for _, current := range arry[2:] {

		// do we have a new highest product of 3?
		// it's either the current highest,
		// or the current times the highest product of two
		// or the current times the lowest product of two
		temp := max(highestProdOf3, current*highestProdOf2)
		highestProdOf3 = max(temp, current*lowestProdOf2)

		// do we have a new highest product of two?
		temp = max(highestProdOf2, current*highest)
		highestProdOf2 = max(temp, current*lowest)

		// do we have a new lowest product of two?
		temp = min(lowestProdOf2, current*highest)
		lowestProdOf2 = min(temp, current*lowest)

		// do we have a new highest?
		highest = max(highest, current)

		// do we have a new lowest?
		lowest = min(lowest, current)
	}

	return highestProdOf3, nil
}
