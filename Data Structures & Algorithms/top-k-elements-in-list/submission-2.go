func topKFrequent(nums []int, k int) []int {
   m1 := make(map[int]int) 
   for _, v := range nums {
    m1[v]++
   }

   m2 := make(map[int]struct{}, k)
   for i := 0; i < k; i++ {
    var cm int
    var ck int
    for k, v := range m1 {
        if _, ok := m2[k]; ok {
           continue 
        }

        if v > cm {
            cm = v
            ck = k
        }
    }

    m2[ck] = struct{}{}
   }

   r := make([]int, 0, k)
   for k := range m2 {
    r = append(r, k) 
   }

   return r
}
