/**
 * @Time : 2020-12-26 21:50
 * @Author : yz
 */

package main

import (
	"fmt"
)

// f[i][j] = min(f[i-1][j], f[i][j-1]) + grid[i][j]

//       j
//	  i	   0  1  2
//		 0 1  3  1
//		 1 1  5  1
//		 2 4  2  1

//       j
//	  i	   0  1  2
//		 0 1  4  5
//		 1 1  6  6
//		 2 5  7  7

func minPathSum(grid [][]int) int {

	// 行数
	m := len(grid)

	// 列数
	n := len(grid[0])

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

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
