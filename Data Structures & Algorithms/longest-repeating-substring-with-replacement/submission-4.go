func characterReplacement(s string, k int) int {
	m := make(map[byte]int)
	var res int
	var l int
	for r := range s {
		m[s[r]]++
		windowSize := r - l + 1

		for {
			var c int
			for _, v := range m {
				c = max(c, v)
			}

			
			if (windowSize-c) <= k {
				break
			}

			m[s[l]]--
			l++
			windowSize = r-l+1
		}

		res = max(res, windowSize)
	}
	return res
}
