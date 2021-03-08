package algcode

import "sort"

/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	if len(nums) < 3 {
		return ret
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := 0 - nums[i]
		l := i + 1
		r := len(nums) - 1
		for l < r {
			if nums[l]+nums[r] == target {
				tmp := []int{nums[i], nums[l], nums[r]}
				ret = append(ret, tmp)
				for l+1 < len(nums)-1 && nums[l] == nums[l+1] {
					l++
				}
				l++
				for r-1 > l && nums[r] == nums[r-1] {
					r--
				}
				r--
			} else if nums[l]+nums[r] > target {
				r--
			} else {
				l++
			}
		}
	}
	return ret
}
