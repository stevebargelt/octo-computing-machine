package apple

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

// GetMaxProfit return the most profit you can make from one buy and one sale during a day
// Did O(n^2) first (two loops)
// the key was being greedy
// A greedy algorithm iterates through the problem space taking the optimal
//		solution "so far," until it reaches the end.
func GetMaxProfit(prices []int) (int, error) {

	if prices == nil {
		return 0, errors.New("prices arrany can not be nil")
	}

	if len(prices) < 2 {
		return 0, errors.New("Need at least two values to compare")
	}
	minPrice := prices[0]
	maxProfit := prices[1] - prices[0]

	for i := 1; i < len(prices); i++ {
		thisNum := prices[i]
		potentialProfit := thisNum - minPrice
		maxProfit = max(maxProfit, potentialProfit)
		minPrice = min(minPrice, thisNum)
	}
	return maxProfit, nil
}
