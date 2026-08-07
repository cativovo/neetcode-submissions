func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strs {
		k := createKey(v)
		m[k] = append(m[k], v)
	}
	s := make([][]string, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

const start = 97 // 'a' code point
const lettersInAlphabet = 26

func createKey(str string) string {
	s := make([]int, lettersInAlphabet)
	for _, v := range str {
		s[v-start]++
	}
	var k strings.Builder
	k.Grow(len(s))
	for _, v := range s {
		k.WriteString(strconv.Itoa(v))
		k.WriteRune('-')
	}
	return k.String()
}