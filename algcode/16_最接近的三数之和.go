/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

*/

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ret := nums[0] + nums[1] + nums[2]
	for i:= 0; i< len(nums)-2; i++{
		tt := target - nums[i]
		l := i+1
		r := len(nums) -1
		for l< r{
			tmp := nums[l] + nums[r]
			if tmp == tt {
				return target
			}
			if abs(tmp-tt) < abs(ret - target){
				ret = tmp + nums[i]
			
			}
			if tmp > tt {
				for r > l  &&  nums[r-1] == nums[r]{
					r--
				}
				r--
			}else {
				for r > l && nums[l] == nums[l+1]{
					l++
				}
				l++
			}
		}
	}
	return ret
}

func abs(a int) int {
	if a < 0{
		a = -1 * a
	}
	return a
}
