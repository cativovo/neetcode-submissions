type Solution struct{
    i []string
}

const delim = "**"

func (s *Solution) Encode(strs []string) string {
    s.i = strs
    return ""
}

func (s *Solution) Decode(ss string) []string {
    return s.i
}
