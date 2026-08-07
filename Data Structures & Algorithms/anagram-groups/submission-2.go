func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)

    for _, v := range strs {
       k := createKey(v) 
       m[k] = append(m[k], v)
    }

    r := make([][]string, 0, len(m))
    for _, v := range m {
       r = append(r, v) 
    }

    return r
}




func createKey(str string) string {
    const aASCIICode = 97
    const lettersInAlphabet = 26
    s := make([]int, 26)

    for _, v := range str {
        s[v-aASCIICode]++
    }

    var k string
    for _, v := range s {
       k += strconv.Itoa(v) + "-"
    }

    return k
}