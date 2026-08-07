type Solution struct{}

const delimiter = '*'

func (s *Solution) Encode(strs []string) string {
	var b strings.Builder
	for _, v := range strs {
		b.WriteString(strconv.Itoa(len(v)))
		b.WriteRune(delimiter)
		b.WriteString(v)
	}
	return b.String()
}

func (s *Solution) Decode(encoded string) []string {
	var result []string
	var i int
	for i < len(encoded) {
		var b strings.Builder
		for {
			v := rune(encoded[i])
			if v == delimiter {
				break
			}
			b.WriteRune(v)
			i++
		}

		c, err := strconv.Atoi(b.String())	
		if err != nil {
			panic(err)	
		}

		start := i+1
		end := start+c
		str := encoded[start:end]	
		result = append(result, str)
		i = end
	}
	return result
}
