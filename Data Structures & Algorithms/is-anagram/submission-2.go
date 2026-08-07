func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
       return false 
    }

    sm := make(map[rune]int)
    for _, v := range s {
       sm[v]++
    }

    tm := make(map[rune]int)
    for _, v := range t {
       tm[v]++
    }

    for k, v := range sm {
        if tm[k] != v {
            return false
        }
    }

    return true
}
