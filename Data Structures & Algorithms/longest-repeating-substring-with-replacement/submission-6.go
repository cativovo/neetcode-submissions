func characterReplacement(s string, k int) int {
	m := make(map[byte]int)
	var res int
	var l int
	for r := range s {
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
		res = max(res, r-l+1)
	}
	return res
}
