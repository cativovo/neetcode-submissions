func maxProfit(prices []int) int {
	var res int
	var l int
	for r := range prices {
		// prices[l] = buy price
		// prices[r] = sell price
		for prices[l] > prices[r] {
			l++
		}
		res = max(res, prices[r]-prices[l])
	}
	return res
}
