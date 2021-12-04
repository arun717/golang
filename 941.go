func validMountainArray(arr []int) bool {
    N := len(arr)
    i := 0
    for(i+1 < N && arr[i] < arr[i+1]) {
        i += 1
    }
    fmt.Println("i start = ",i)
    if i == 0 ||  i == N-1 {
        return false
    }
    for i+1 < N && arr[i] > arr[i+1] {
        i += 1
    }
    fmt.Println("i end = ",i)
    return i == N-1   
}