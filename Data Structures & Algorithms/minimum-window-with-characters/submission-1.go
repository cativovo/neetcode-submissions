func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	
	want := make(map[byte]int, len(t))
	for i := range t {
		want[t[i]]++
	}

	wantLen := len(want)
	var haveLen int


	var res [2]int
	resLen := math.MaxInt
	window := make(map[byte]int)
	var l int
	for r := range s {
		k := s[r]	
		window[k]++
		if _, ok := want[k]; ok && window[k] == want[k]{
			haveLen++	
		}

		for haveLen == wantLen {
			windowSize := r-l+1
			if windowSize < resLen {
				res = [2]int{l, r}
				resLen = windowSize
			}	

			k := s[l]
			window[k]--
			if _, ok := want[k]; ok && window[k] < want[k] {
				haveLen--
			}
			l++
		}
	}

	if resLen == math.MaxInt {
		return ""
	}
	return s[res[0]:res[1]+1]
}
