func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	want := make(map[byte]int)
	for i := range len(t) {
		want[t[i]]++
	}
	wantLen := len(want)

	window := make(map[byte]int)
	resLen := math.MaxInt
	var haveLen int
	var res [2]int

	var l int
	for r := range len(s) {
		k := s[r]
		window[k]++
		if v, ok := window[k]; ok && v == want[k] {
			haveLen++
		}

		for wantLen == haveLen {
			if (r - l + 1) < resLen {
				res = [2]int{l, r + 1}
				resLen = r-l+1
			}

			k := s[l]
			window[k]--
			if v, ok := window[k]; ok && v < want[k] {
				haveLen--
			}
			l++
		}

	}

	if resLen == math.MaxInt {
		return ""
	}

	return s[res[0]:res[1]]
}
