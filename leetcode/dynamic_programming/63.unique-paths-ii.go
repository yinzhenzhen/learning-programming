/**
 * @Time : 2020-12-26 15:59
 * @Author : yz
 */

package dynamic_programming

/**
标签：数组、动态规划、矩阵
难度：中等

题目内容：
63.不同路径II
给定一个 m*n 的整数数组 grid。一个机器人初始位于左上角（即 grid[0][0]）。
机器人尝试移动到右下角（即 grid[m - 1][n - 1]）。机器人每次只能向下或者向右移动一步。
网格中的障碍物和空位置分别用 1 和 0 来表示。机器人的移动路径中不能包含任何有障碍物的方格。
返回机器人能够到达右下角的不同路径数量。
测试用例保证答案小于等于 2 * 10的9次方。

解题思路：见代码注释及状态方程
f[i][j] = f[i-1][j] + f[i][j-1]   obstacleGrid[i][j]=0
f[i][j] = 0                       obstacleGrid[i][j]=1

	         j
		  i	   0  1  2  3  4  5
			 0 1  1  1  1  1  1
			 1 1  √  1  √  1  2
			 2 1  1  2  2  3  5
			 3 1  2  4  √  3  8
			 4 1  3  7  7  10 18
			 5 1  4  11 18 28 46
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	// 行数
	m := len(obstacleGrid)

	// 列数
	n := len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	// 初始化二维数组
	var f [][]int
	for i := 0; i < m; i++ {
		temp := make([]int, len(obstacleGrid[i]))
		f = append(f, temp)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果没有障碍物
			if obstacleGrid[i][j] == 0 {
				if i == 0 && j == 0 {
					// 如果位置索引都为 0，就是第一次经过的地方，则置为1
					f[i][j] = 1
				} else if i == 0 {
					// 如果位置索引 i 为 0，则代表第一横排，第一横排当前索引能到达的路线和这一横排的前一个位置一致，都为1，因为只能向下，不能向上
					f[i][j] = f[i][j-1]
				} else if j == 0 {
					// 如果位置索引 j 为 0，则代表第一竖排，第一竖排当前索引能到达的路线和这一竖排的上一个位置一致，都为1，因为只能向右，不能向左
					f[i][j] = f[i-1][j]
				} else {
					// 其他位置则是根据最近的左边和上边之和来确定有多少条路线的
					f[i][j] = f[i-1][j] + f[i][j-1]
				}
			}
		}
	}

	return f[m-1][n-1]
}
