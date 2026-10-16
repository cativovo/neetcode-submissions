func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]struct{})
	var res int
	var l int
	for r := range s {
		_, ok := m[s[r]]
		if ok {
				// ipasok lang yung s[r] sa map pag naalis na yung duplicate
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
