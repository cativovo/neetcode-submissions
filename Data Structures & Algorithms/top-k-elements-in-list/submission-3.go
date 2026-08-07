func topKFrequent(nums []int, k int) []int {
   m := make(map[int]int) 
   for _, v := range nums {
    m[v]++
   }

   s := make([][]int, len(nums)+1)
   for k, v := range m {
    s[v] = append(s[v], k)
   }

   r := make([]int, 0)
   for i := len(s) - 1; i >= 0; i--{
    r = append(r, s[i]...)
   }

    rr := make([]int, 0, k)
    for i := 0; i < k; i++ {
        rr = append(rr, r[i])
    }
   return rr
}
