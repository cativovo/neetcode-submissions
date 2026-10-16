func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]struct{})
	var res int
	var l int
	var r int
	for r < len(s) {
		_, ok := m[s[r]]
		if ok {
			for {
				_, ok := m[s[r]]
				if !ok {
					break
				}
				delete(m, s[l])
				l++
			}
		}

		m[s[r]] = struct{}{}

		res = max(res, len(m))
		r++
	}
	return res
}
