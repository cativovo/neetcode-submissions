func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		m[v] = struct{}{}
	}
	
	var result int

	for _, v := range nums {
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
