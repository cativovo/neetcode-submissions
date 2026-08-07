func hasDuplicate(nums []int) bool {
   s := make(map[int]struct{})

   for _, v := range nums {
    if _, ok := s[v]; ok {
       return true 
    }
    s[v] = struct{}{}
   }

   return false
}
