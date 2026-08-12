func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		m[v] = struct{}{}
	}

	var result int
	for _, v := range nums {
		// skip pag hindi siya start ng sequence
		// for example: [1,2,6,3]
		// mag count lang ng sequence pag yung v = 1 or 6
		// skip yung 2 at 3 kasi maka-count sila pag v = 1
		v--
		_, ok := m[v]
		if ok {
			continue
		}

		var c int
		for {
			v++
			_, ok := m[v]
			if !ok {
				break
			}
			c++
		}
		result = max(result, c)
	}
	return result
}
