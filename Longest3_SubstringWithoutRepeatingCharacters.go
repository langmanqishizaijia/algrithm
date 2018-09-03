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
	pre    := -1
	mp     := make(map[byte]int)

  
	for i:= 0; i< len(s) ; i++ {

		if _, ok := mp[s[i]] ; ok {

			if mp[s[i]] >= pre {
				pre = mp[s[i]]
			}

		}
		if maxlen < i-pre  {
		   maxlen = i-pre
		}
		mp[s[i]] = i
	}

 
    
    return maxlen
}
