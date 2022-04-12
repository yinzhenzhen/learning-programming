/**
 * @Time : 2020-12-26 15:59
 * @Author : yz
 */

package main

import "fmt"

// f[i][j] = f[i-1][j] + f[i][j-1]   obstacleGrid[i][j]=0
// f[i][j] = 0                       obstacleGrid[i][j]=1

//       j
//	  i	   0  1  2  3  4  5
//		 0 1  1  1  1  1  1
//		 1 1  √  1  √  1  2
//		 2 1  1  2  2  3  5
//		 3 1  2  4  √  3  8
//		 4 1  3  7  7  10 18
//		 5 1  4  11 18 28 46
func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	// 行数
	m := len(obstacleGrid)

	// 列数
	n := len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	var f [][]int

	for i := 0; i < m; i++ {
		temp := make([]int, len(obstacleGrid[i]))
		f = append(f, temp)
	}

	for i := 0; i < m; i++ {

		for j := 0; j < n; j++ {

			if obstacleGrid[i][j] == 0 {
				if i == 0 && j == 0 {
					f[i][j] = 1
				} else if i == 0 {
					f[i][j] = f[i][j-1]
				} else if j == 0 {
					f[i][j] = f[i-1][j]
				} else {
					f[i][j] = f[i-1][j] + f[i][j-1]
				}

			}

		}
	}

	return f[m-1][n-1]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 1}}))

	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0},
		{1, 1},
		{0, 0},
	}))

	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}))
}
