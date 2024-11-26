package easy

func maxProfit(prices []int) int {
    minBuyPrice := prices[0]
    maxProfit := 0
    for i := 1; i < len(prices); i++ {
        if prices[i] - minBuyPrice > maxProfit {
            maxProfit = prices[i] - minBuyPrice
        }
        if prices[i] < minBuyPrice {
            minBuyPrice = prices[i]
        }
    }
    return maxProfit 
}