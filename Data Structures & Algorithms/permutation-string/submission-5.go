func checkInclusion(s1 string, s2 string) bool {
	m := make(map[byte]int, len(s1))
	for i := range len(s1) {
		m[s1[i]]++
	}

	w := make(map[byte]int)
	var l int
	for r := range len(s2) {
		w[s2[r]]++

		if r-l+1 == len(s1) {
			t := true
			for k, v := range m {
				vv := w[k]
				if vv != v {
					t = false
					break
				}
			}
			if t {
				return true
			}

			w[s2[l]]--
			l++
		}
	}
	return false
}
