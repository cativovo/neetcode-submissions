func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
       return false 
    }

    sm := make(map[rune]int)
    tm := make(map[rune]int)
    tr := []rune(t)
    for i, v := range s {
       sm[v]++
       tm[tr[i]]++
    }

    for k, v := range sm {
        if tm[k] != v {
            return false
        }
    }

    return true
}
