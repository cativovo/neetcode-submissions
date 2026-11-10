func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]struct{})
	var result int
	var l int
	for r := range len(s) {
		rk := s[r]

		for {
			_, ok := m[rk]
			if !ok {
				break
			}
			delete(m, s[l])
			l++
		}

		m[rk] = struct{}{}
		result = max(result, r-l+1)
	}
	return result
}