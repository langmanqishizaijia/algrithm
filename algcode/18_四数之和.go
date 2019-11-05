/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

*/


var res [][]int
func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    //fmt.Printf("%v",nums)
    if len(nums) < 4 {
	  return [][]int{}
    }
   res = make([][]int, 0)
   tmp := make([]int, 0)
   calDFS(0, 0 ,nums, tmp,target)

	return res
}
func calDFS(i int, cur int ,nums []int, tmp []int ,target int ){
	if len(tmp) == 4{
		if cur == target{
			kk := make([]int, 4)
			for k, v :=range tmp{
				kk[k] = v
			}
			res = append(res,kk)
		}
		return
	}else if len(tmp) > 4{
		return
	}

	for j:= i ; j< len(nums); j++ {
		if j > i && nums[j-1] == nums[j] {
			continue
		}
		calDFS(j+1, cur + nums[j], nums, append(tmp,nums[j]),target)


	}
}
