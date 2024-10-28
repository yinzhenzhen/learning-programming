package dynamic

import (
	"fmt"
)

/**
标签：字符串、动态规划
难度：中等

题目内容：
HJ52.计算字符串的编辑距离
Levenshtein 距离，又称编辑距离，指的是两个字符串之间，由一个转换成另一个所需的最少编辑操作次数。
许可的编辑操作包括将一个字符替换成另一个字符，插入一个字符，删除一个字符。
编辑距离的算法是首先由俄国科学家 Levenshtein 提出的，故又叫 Levenshtein Distance 。
例如：
字符串A: abcdefg
字符串B: abcdef
通过增加或是删掉字符 ”g” 的方式达到目的。这两种方案都需要一次操作。把这个操作所需要的次数定义为两个字符串的距离。
要求：
给定任意两个字符串，写出一个算法计算它们的编辑距离。

解题思路：from题解
这题考的是levenshtein距离的计算，需要运用动态规划去解决该类问题。
传递公式：
lev[i][j]用来表示字符串a的[1...i]和字符串b[1...j]的levenshtein距离；
插入和删除操作互为逆过程：a删除指定字符变b等同于b插入指定字符变a；
如果a[i] == b[j]，则说明a[i]和b[j]分别加入a，b之后不会影响levenshtein距离，lev[i][j] = lev[i-1][j-1] + 0;
如果a[i] != b[j]，则需要考虑3种情况的可能：
a中插入字符，即lev[i][j] = lev[i-1][j] + 1;
b中插入字符，即lev[i][j] = lev[i][j-1] + 1;
a[i]替换成b[j]，lev[i][j] = lev[i-1][j-1] + 1;
取这4种情况的最小值。
动态规划 dp[i][j] = min(dp[i-1][j]+1,dp[i][j-1]+1,dp[i-1][j-1]+1 {是否加1视情况而定})

*/

func levenshtein() {
	a := ""
	b := ""
	for {
		n, _ := fmt.Scan(&a, &b)
		if n == 0 {
			break
		} else {
			c := lev(a, b)
			fmt.Printf("%d\n", c)
		}
	}
}

func lev(a, b string) int {
	m := len(a)
	n := len(b)
	// 初始化
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 边界条件初始化
	for i := 0; i < m+1; i++ {
		dp[i][0] = i // word1[i] 变成 word2[0], 删掉 word1[i], 需要 i 部操作
	}
	for j := 0; j < n+1; j++ {
		dp[0][j] = j // word1[0] 变成 word2[j], 插入 word1[j]，需要 j 部操作
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else { // Min(插入，删除，替换)
				dp[i][j] = Min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

func Min(args ...int) int {
	min := args[0]
	for _, item := range args {
		if item < min {
			min = item
		}
	}
	return min
}
