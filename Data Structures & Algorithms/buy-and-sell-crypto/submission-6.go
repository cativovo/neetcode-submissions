func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	l := 0
	r := 1
	p := 0
	for r < len(prices) {
		lp := prices[l]
		rp := prices[r]
		p = max(p, rp-lp)
		if lp > rp {
			l = r
		}
		r++
	}
	return p
}
