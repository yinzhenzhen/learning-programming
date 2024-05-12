/**
 * @Time : 2020-12-26 13:01
 * @Author : yz
 */

package dynamic_programming

//       j
//	  i	   0  1  2  3  4  5
//		 0 √  a  b  d  g  k
//		 1    √  c  e  h  l
//		 2       √  f  i  m
//		 3          √  j  n
//		 4             √  o
//		 5                √

// 状态方程：dp[i][j] = s[i]==s[j] && dp[i+1][j-1]
func longestPalindrome(s string) string {

	length := len(s)
	if length < 2 {
		return s
	}

	maxLen := 1
	begin := 0
	end := 0

	var dp [][]bool

	for i := 0; i < len(s); i++ {
		temp := make([]bool, length)
		dp = append(dp, temp)
		dp[i][i] = true
	}

	ss := []byte(s)

	for j := 1; j < len(s); j++ {

		for i := 0; i < j; i++ {

			if ss[i] != ss[j] {
				dp[i][j] = false
				continue
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if j-i+1 > maxLen && dp[i][j] {
				maxLen = j - i + 1
				begin = i
				end = j + 1
			}
		}
	}

	if maxLen == 1 {
		return string(ss[0:1])
	}

	ss = ss[begin:end]

	return string(ss)
}
