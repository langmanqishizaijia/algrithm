func missingNumber(nums []int) int {
    num := len (nums)
    m := make(map[int]int, 0)
    for _, v := range nums {
        m[v] = v
    }
    
    for i:= 0; i< num ;i++ {
        if _, ok := m[i];!ok{
            return i
            break
        }
    }
    return num
}
