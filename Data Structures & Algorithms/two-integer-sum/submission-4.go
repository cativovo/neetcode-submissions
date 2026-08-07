func twoSum(nums []int, target int) []int {
   m := make(map[int]int) 
   for i, v := range nums {
        m[target - v] = i
   }

   for i, num := range nums {
    if v, ok := m[num]; ok && v != i{
        a := min(v, i)
        b := max(v, i)
        return []int{a, b}
    }
   }

   return nil
}
