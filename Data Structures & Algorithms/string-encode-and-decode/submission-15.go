type Solution struct{}

const delim = "*"

func (s *Solution) Encode(strs []string) string {
	var r string
	for _, v := range strs {
		r += strconv.Itoa(countRunes(v)) + delim + v
	}
	return r
}

func (s *Solution) Decode(ss string) []string {
	fmt.Println(ss)
	var r []string
	rr := []rune(ss)
	ssLen := countRunes(ss)
	for i := 0; i < ssLen; i++ {
		v := rr[i]
		if isNum(v) {
			var di int
			for j := i + 1; j < ssLen; j++ {
				v := string(rr[j])
				if v == delim {
					di = j
					break
				}
			}

			numStr := string(rr[i:di])
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}

			var str string
			for k := range num {
				str += string(rr[i+k+len(numStr)+1])
			}

			r = append(r, str)
			i += num + 1
		}
	}
	return r
}

var m = map[rune]struct{}{
	'0': {},
	'1': {},
	'2': {},
	'3': {},
	'4': {},
	'5': {},
	'6': {},
	'7': {},
	'8': {},
	'9': {},
}

func isNum(r rune) bool {
	_, ok := m[r]
	return ok
}

func countRunes(s string) int {
	var r int
	for range s {
		r++
	}
	return r
}
