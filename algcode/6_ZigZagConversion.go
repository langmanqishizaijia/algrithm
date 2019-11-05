/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
*/

func convert(s string, numRows int) string {
   
	
	l := len(s)
	if numRows == 1 {
		return s
	}
	mp := make([][]byte,numRows)
	for i:= 0; i< numRows;i++{
		mp[i] = make([]byte, l)
	}


	a := 0
	b := 0
	dir := 0
	for i:= 0; i <l ;i++{
		mp[a][b] = s[i]
		if dir == 0{
			a++
			if a == numRows{
				dir = 1
				a -= 2
				b++
			}
		}else {
			a--
			b++
			if a== -1{
				a += 2
				dir = 0
			}
		}
	}

	ret := ""
	for i := 0 ; i< numRows; i++ {
		for j:= 0; j< l; j++ {
			if mp[i][j]  != 0 {
				ret = ret + string(mp[i][j])
			}
		}
	}
	return ret
}
