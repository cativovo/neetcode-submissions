func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	m := make(map[byte]int)
	for i := range len(t) {
		m[t[i]]++
	}

	w := make(map[byte]int)
	var res string
	var l int
	for r := range len(s) {
		w[s[r]]++

		for {
			valid := true
			for k, v := range m {
				vv := w[k]
				if vv < v {
					valid = false
					break
				}
			}

			if !valid {
				break
			}

			if res == "" || (r-l+1) < len(res) {
				res = s[l:r+1]
			}

			w[s[l]]--
			l++
		}

	}

	return res
}