/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
 

*/


func twoSum(nums []int, target int) []int {
    
    m := make(map[int]int, 0)
    res := make([]int, 0)
    for k, v := range nums {
        m[v] = k
    }
    for i:= 0; i< len(nums); i++{
        t := target-nums[i]
        if _, ok := m[t]; ok && m[t] != i {
            return append(res, i, m[t])
        }
    }
    return res
}
