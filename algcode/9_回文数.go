package algcode

import "strconv"

/*

判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
进阶:

你能不将整数转为字符串来解决这个问题吗？


*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)
	return judge(str)
}

func judge(str string) bool {
	if len(str) == 1 {
		return true
	}
	i := 0
	j := len(str) - 1

	for i < j {

		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

//另外一种方法
/*
将原数做反转，然后新老比较
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}
	o := x
	y := 0
	for x != 0 {
		tmp := x % 10
		y = y*10 + tmp
		x = x / 10
	}
	return o == y
}
