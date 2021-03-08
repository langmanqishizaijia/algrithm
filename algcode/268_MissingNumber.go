package algcode

/*
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

Example 1:

Input: [3,0,1]
Output: 2
Example 2:

Input: [9,6,4,2,3,5,7,0,1]
Output: 8
Note:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?
*/

func missingNumber(nums []int) int {
	num := len(nums)
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v] = v
	}

	for i := 0; i < num; i++ {
		if _, ok := m[i]; !ok {
			return i
			break
		}
	}
	return num
}
