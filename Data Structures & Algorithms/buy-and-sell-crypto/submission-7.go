func maxProfit(prices []int) int {
	var res int
	var l int
	for r := range prices {
		for prices[l] > prices[r] {
			l++
		}
		res = max(res, prices[r]-prices[l])
	}
	return res
}
