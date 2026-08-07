type Solution struct{}

const delim byte = '*'

func (s *Solution) Encode(strs []string) string {
    var r string
    for _, v := range strs {
        r += strconv.Itoa(len(v)) + string(delim) + v
    }
    return r
}

func (s *Solution) Decode(ss string) []string {
    var r []string
    for i := 0; i < len(ss); {
       j := i 
       for ss[j] != delim {
        j++
       }

       length, err := strconv.Atoi(ss[i:j])
       if err != nil {
        panic(err)
       }

       i = j+1

       word := ss[i:i+length]
       r = append(r, word)
       i += length 
    }
    return r
}
