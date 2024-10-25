/**
 * @Time : 2020-12-26 21:50
 * @Author : yz
 */

package dynamic_programming

/**
标签：数组、动态规划、矩阵
难度：中等

题目内容：
64.最小路径和
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。

解题思路：见代码注释及状态方程
f[i][j] = min(f[i-1][j], f[i][j-1]) + grid[i][j]

         j
	  i	   0  1  2
		 0 1  3  1
		 1 1  5  1
		 2 4  2  1

         j
	  i	   0  1  2
		 0 1  4  5
		 1 1  6  6
		 2 5  7  7
*/

func minPathSum(grid [][]int) int {

	// 行数
	m := len(grid)

	// 列数
	n := len(grid[0])

	// 初始化二维数组
	var f [][]int
	for i := 0; i < m; i++ {
		temp := make([]int, len(grid[i]))
		f = append(f, temp)
	}

	for i := 0; i < m; i++ {

		for j := 0; j < n; j++ {

			if i == 0 && j == 0 {
				f[0][0] = grid[0][0]
			} else if i == 0 {
				f[i][j] = f[i][j-1] + grid[i][j]
			} else if j == 0 {
				f[i][j] = f[i-1][j] + grid[i][j]
			} else {
				f[i][j] = MinInt(f[i][j-1], f[i-1][j]) + grid[i][j]
			}
		}
	}

	return f[m-1][n-1]
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
