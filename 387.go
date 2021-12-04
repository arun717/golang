func firstUniqChar(s string) int {
    m := make(map[string]int)
    str := s
    for _, r := range str {
        c := string(r)
        fmt.Println(c)
        if m[c] != 0{
            m[c]+=1
        }else{
            m[c] = 1
        }
    }
    fmt.Println(m)
    for i, r := range str {
        if(m[string(r)]==1){
            return i
        }
    }
    return -1
}