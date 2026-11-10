func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]struct{})
	var res int
	var l int
	for r := range s {
		for {
			_, ok := m[s[r]]	
			if !ok {
				break
			}
			delete(m, s[l])
			l++
		}
		m[s[r]] = struct{}{}
		res = max(res, len(m))
	}
	return res
}
