func shuffle(nums []int, n int) []int {
    s1 := make([]int, 0)
    for i:=0; i<n; i++ {
        s1 = append(s1,nums[i])
        s1 = append(s1,nums[n+i])    
    }
    return s1
}