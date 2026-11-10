func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	var res int
	var l int
	for r := range s {
		for {
			_, ok := m[s[r]]
			if !ok {
				break
			}
			k := s[l]
			m[k]--
			if m[k] == 0 {
				delete(m, k)
			}
			l++
		}
		m[s[r]]++
		res = max(res, len(m))
	}
	return res
}
