package algcode

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。


*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	pre := strs[0]
	for i := 1; i < len(strs); i++ {
		pre = getComenPre(strs[i], pre)
	}
	return pre
}
func getComenPre(s1, s2 string) string {
	ret := ""
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			return ret
		} else {
			ret += string(s1[i])
		}
	}

	return ret
}
