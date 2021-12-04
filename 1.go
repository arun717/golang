func twoSum(nums []int, target int) []int {
    var res []int
    var i,j int
    for i=0;i<len(nums)-1;i++ {
        for j=i+1;j<len(nums);j++ {
            if nums[i]+nums[j] == target {
                res = append(res,i)
                res = append(res,j)
                break
            }
        }
    }
    return res
}