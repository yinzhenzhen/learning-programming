/**
 * @Time : 2020-11-30 15:59
 * @Author : yz
 */

package swordfinger

/*
*
二维数组中的查找
*/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	// 一行多少元素
	m := len(matrix[0])
	// 一列多少元素
	n := len(matrix)

	// 与右上角元素进行比较，如果比右上角大，往下查找，比右上角小，往左查找
	for i, j := 0, m-1; i < n && j >= 0; {

		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}

	return false
}
