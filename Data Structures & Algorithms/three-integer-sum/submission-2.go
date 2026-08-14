func threeSum(nums []int) [][]int {
	const target = 0
	sort.Ints(nums)	
	var result [][]int
	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}

		l := i+1
		r := len(nums)-1
		for l < r {
			a := nums[l]
			b := nums[r]
			sum := v+a+b
			if sum < target {
				l++
			} else if sum > target {
				r--
			} else {
				result = append(result, []int{v,a,b})
				l++
				for  nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}
	return result
}
