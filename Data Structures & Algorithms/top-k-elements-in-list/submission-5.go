func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	s := make([]int, 0, k)
	for {
		var kk int
		var vv int
		for k, v := range m {
			if v > vv {
				kk = k
				vv = v
			}
		}
		s = append(s, kk) 
		if len(s) == k {
			return s
		}

		delete(m, kk)
	}
}