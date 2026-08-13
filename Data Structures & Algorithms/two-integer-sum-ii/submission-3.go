func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers)-1
	for l < r {
		// sorted yung numbers, mas maliit lagi yung left side
		sum := numbers[r] + numbers[l]
		if sum == target {
			return []int{l+1, r+1}
		}

		if sum > target {
			r--
		} else {
			l++
		} 

	}
	return nil
}
