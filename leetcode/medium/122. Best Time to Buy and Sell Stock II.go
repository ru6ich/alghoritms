package medium

import "math"

func maxProfit(prices []int) int {
	minBuyPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-minBuyPrice > 0 {
			maxProfit += prices[i] - minBuyPrice
			minBuyPrice = math.MaxInt64
		}
			if prices[i] < minBuyPrice {
			minBuyPrice = prices[i]
		}
	}
	return maxProfit
}