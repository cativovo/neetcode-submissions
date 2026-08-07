func twoSum(nums []int, target int) []int {
   m := make(map[int]int) 
   for i, v := range nums {
	vv, ok := m[v]
	if ok {
		return []int{vv, i}
	}

	m[target-v]=i
   }
   return nil
}
