/**
 * @Time : 2020-12-26 13:01
 * @Author : yz
 */

package dynamic_programming

/**
标签：双指针、字符串、动态规划
难度：中等

题目内容：
5.最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。

解题思路：见代码注释及状态方程
图1：
         j
	  i	   0  1  2  3  4  5
		 0 √  a  b  d  g  k
		 1    √  c  e  h  l
		 2       √  f  i  m
		 3          √  j  n
		 4             √  o
		 5                √
状态方程：dp[i][j] = s[i]==s[j] && dp[i+1][j-1]
*/

func longestPalindrome(s string) string {

	length := len(s)
	if length < 2 {
		return s
	}

	maxLen := 1
	begin := 0
	end := 0

	// 构建一个大小为s的长度的布尔值二维数组，斜对角给默认为true，因为斜对角就是本身
	var dp [][]bool
	for i := 0; i < len(s); i++ {
		temp := make([]bool, length)
		dp = append(dp, temp)
		dp[i][i] = true
	}

	ss := []byte(s)

	// i为左端索引，j为右端索引
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if ss[i] != ss[j] {
				// 如果i和j所在位置的字符不相等，标记为false，并进行下一次标记
				dp[i][j] = false
				continue
			} else {
				// 如果i和j所在位置的字符相等
				if j-i < 3 {
					// 如果两个字符串是相邻或者中间只隔了一个字符，则标记为true
					dp[i][j] = true
				} else {
					// 如果两个字符串相隔了大于1个字符的距离，则需要根据 i+1 和 j-1 这两个位置来判断是否相等
					dp[i][j] = dp[i+1][j-1]
				}
			}
			// 如果本次回文字符串长度大于之前的最大值，则记录首尾字符所在索引位置以及回文长度
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
