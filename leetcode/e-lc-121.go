package main

func maxProfit(prices []int) int {
	profit := 0
	stockToBuy := prices[0]
	for i := 1; i < len(prices); i++ {
		if stockToBuy > prices[i] {
			stockToBuy = prices[i]
		}

		currentProfit := prices[i] - stockToBuy

		if currentProfit > profit {
			profit = currentProfit
		}
	}

	return profit
}
