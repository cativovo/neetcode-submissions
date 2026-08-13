func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	r := []rune(s)
	i := 0
	j := len(r)-1
	for i < j {
		if !isAlphanumeric(r[i]) {
			i++
			continue
		}

		if !isAlphanumeric(r[j]) {
			j--
			continue
		}

		if r[i] != r[j] {
			fmt.Println(string(r[i]), string(r[j]))
			return false
		}

		i++
		j--
	}
	return true
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}