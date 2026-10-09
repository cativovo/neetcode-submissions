func maxArea(heights []int) int {
	l := 0
	r := len(heights)-1
	m := 0
	for l < r {
		lh := heights[l]
		rh := heights[r]
		a := min(lh, rh) * (r - l)
	    m = max(m, a)
		
		if lh < rh {
			l++
		} else {
			r--
		}
	}
	return m
}
