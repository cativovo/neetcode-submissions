func checkInclusion(s1 string, s2 string) bool {
	size := len(s1)
	var l int
	
	m1 := make(map[byte]int, size)
	for i := range s1 {
		m1[s1[i]]++
	}

	m2 := make(map[byte]int)

	R:
	for r := range s2 {
		m2[s2[r]]++
		for (r - l + 1) > size {
			m2[s2[l]]--
			l++
		}
		
		if (r - l + 1) == size {
			for k, v := range m1 {
				vv := m2[k]
				if v != vv {
					continue R
				}
			}

			return true
		}
	}

	return false
}