func productExceptSelf(nums []int) []int {
    r := make([]int, 0, len(nums))
    for i1 := range nums {
        n := 1
       for i2, v2 := range nums {
        if i2 == i1 {
            continue
        }
        n *= v2
       } 
       r = append(r, n)
    }
    return r
}
