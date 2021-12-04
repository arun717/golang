func maximumWealth(accounts [][]int) int {
    var temp int = 0
    var maxi int = 0
    for i:= 0; i< len(accounts); i++{
        temp = 0
        for j:= 0; j< len(accounts[0]); j++{
            temp += accounts[i][j]
            fmt.Println("working on " , accounts[i][j])
        }
        fmt.Println(temp)
        if temp>maxi {
            maxi = temp
        }
    }
    return maxi
}