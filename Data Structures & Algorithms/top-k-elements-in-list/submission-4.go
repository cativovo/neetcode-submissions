func topKFrequent(nums []int, k int) []int {
   m := make(map[int]int) 
   for _, v := range nums {
    m[v]++
   }

   s := make([][]int, len(nums)+1)
   for k, v := range m {
    s[v] = append(s[v], k)
   }

   r := make([]int, 0, k)
   for i := len(s) - 1; i >= 0; i--{
    for _, v := range s[i] {
        r = append(r, v)
        if len(r) == k {
            return r
        }
    }
   }

   return r
}
