
/*

给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/
var ret []string
func generateParenthesis(n int) []string {

	ret = make([]string, 0)
	travelDFS(n,0,0,"")
	return ret
}

func travelDFS (n , left, right int , str string) {

	if  n == right  && left == n   {
		ret = append(ret, str)
		return
	}

	if left < right {
		return
	}

	if left > n {
		return
	}

	if right > n{
		return
	}
	travelDFS(n, left+1, right, str + "(" )
	travelDFS(n, left, right +1 , str + ")")


}
