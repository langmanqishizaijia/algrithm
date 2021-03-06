package algcode

/*
给定一个字符串 (s) 和一个字符模式 (p)。实现支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符。
'*' 匹配零个或多个前面的元素。
匹配应该覆盖整个字符串 (s) ，而不是部分字符串。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释: '*' 代表可匹配零个或多个前面的元素, 即可以匹配 'a' 。因此, 重复 'a' 一次, 字符串可变为 "aa"。
示例 3:

输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个('*')任意字符('.')。
示例 4:

输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 'c' 可以不被重复, 'a' 可以被重复一次。因此可以匹配字符串 "aab"。
示例 5:

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false
*/

func isMatch(s string, p string) bool {

	if len(s) != 0 && len(p) == 0 {
		return false
	}
	//如果s p 都为空 返回true
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	//如果p只有一个的情况 要么是. 要么是字母
	if len(p) == 1 {
		return len(s) == 1 && (s[0] == p[0] || p[0] == '.')
	}

	//若p的第二个字符不为*,若此时s为空则返回false,否则判断首字符是否匹配，且从各自的第二个字符开始调用
	//递归函数匹配
	if p[1] != '*' {
		if len(s) == 0 {
			return false
		}
		return (s[0] == p[0] || p[0] == '.') && isMatch(s[1:], p[1:])
	}

	//若p的第二个字符为*，若s不为空且字符匹配， 调用递归函数匹配s和去掉前两个字符的p,若匹配返回true，否则去掉s的
	//首字母
	for len(s) != 0 && (s[0] == p[0] || p[0] == '.') {
		if isMatch(s, p[2:]) {
			return true
		}
		s = s[1:]
	}

	//返回调用递归函数匹配s和去掉前两个字符的p的结果

	return isMatch(s, p[2:])
}
