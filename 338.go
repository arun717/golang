func countBits(n int) []int {
    var res []int
    for val:=0;val<=n;val++ {
        output := strconv.FormatInt(int64(val), 2)  
        countedOnes := strings.Count(output, "1")
        res = append(res,countedOnes)
    }
    return res
}
