package algcode

/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", which the length is 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring(s string) int {

	maxlen := 0
	pre := -1
	mp := make(map[byte]int)

	for i := 0; i < len(s); i++ {

		if _, ok := mp[s[i]]; ok {

			if mp[s[i]] >= pre {
				pre = mp[s[i]]
			}

		}
		if maxlen < i-pre {
			maxlen = i - pre
		}
		mp[s[i]] = i
	}

	return maxlen
}

/*2021060923*/
func lengthOfLongestSubstring2(s string) int {
	mp := make(map[byte]int)
	if len(s) == 0 {
		return 0
	}
	head := 0
	res := 1
	for i := 0; i < len(s); i++ {
		cur := s[i]
		if loc, ok := mp[cur]; ok {
			if loc >= head {
				head = loc + 1
			}
		}
		mp[cur] = i
		res = max(res, i-head+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
