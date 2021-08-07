package besttimetobuyandsellstock

func maxProfit(prices []int) int {

	minVal := prices[0]
	var currentIdx, maxProfit int

	for currentIdx != len(prices) {

		if prices[currentIdx] < minVal {
			minVal = prices[currentIdx]
		}

		if prices[currentIdx]-minVal > maxProfit {
			maxProfit = prices[currentIdx] - minVal
		}
		currentIdx++

	}

	return maxProfit
}
