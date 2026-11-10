func characterReplacement(s string, k int) int {
	m := make(map[byte]int)
	var result int
	var l int
	for r := range len(s) {
		m[s[r]]++

		for {
			var c int
			for _, v := range m {
				c = max(c, v)
			}

			if ((r-l+1)-c) <= k {
				break
			}

			m[s[l]]--
			l++
		}

		result = max(result, r-l+1)	
	}
	return result
}
