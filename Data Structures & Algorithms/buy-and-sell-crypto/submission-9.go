func maxProfit(prices []int) int {
	var result int
	var l int
	for r := 0; r < len(prices); r++ {
		result = max(result, prices[r]-prices[l])
		for prices[l] > prices[r] {
			l++
		}
	}
	return result
}
