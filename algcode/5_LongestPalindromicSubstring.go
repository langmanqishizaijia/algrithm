package algcode

/*Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
*/
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	maxstr := s[0:1]
	for l := 2; l <= len(s); l++ {
		for i := 0; i+l-1 < len(s); i++ {
			j := i + l - 1
			if s[i] == s[j] && j-i > 1 && dp[i+1][j-1] == 1 {
				dp[i][j] = 1
				if len(maxstr) < l {
					maxstr = s[i : j+1]
				}
			} else if s[i] == s[j] && j-i == 1 {
				dp[i][j] = 1
				if len(maxstr) < l {
					maxstr = s[i : j+1]
				}
			}
		}
	}

	return maxstr
}
