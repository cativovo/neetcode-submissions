func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers)-1
	for i < j {
		a := max(numbers[i], numbers[j]) 
		b := min(numbers[i], numbers[j]) 
		sum := a+b
		if sum == target {
			return []int{i+1, j+1}
		}

		if sum > target {
			j--
		} else {
			i++
		} 

	}
	return nil
}
