/**
 * @Time : 2022/6/29 10:39 上午
 * @Author : yz
 */

package backtracking

import "fmt"

/**
题目内容：
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。

解题思路：

*/

func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	generateCombinations(n, k, 1, c, &res)
	return res
}

func generateCombinations(n, k, start int, c []int, res *[][]int) {
	fmt.Println("enter---")
	if len(c) == k {
		b := make([]int, len(c))
		copy(b, c)
		fmt.Println("over-----", b)
		*res = append(*res, b)
		return
	}
	// i will at most be n - (k - c.size()) + 1
	for i := start; i <= n-(k-len(c))+1; i++ {
		c = append(c, i)
		fmt.Println("-----1", c)
		fmt.Println("----------before", i)
		generateCombinations(n, k, i+1, c, res)
		c = c[:len(c)-1]
		fmt.Println("-----2", c)
		fmt.Println("----------after", i)
	}
	fmt.Println("return")
	return
}
