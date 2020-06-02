/**
 * @Time : 2020-06-02 23:10
 * @Author : yz
 */

package main

import "fmt"

// 超出时间限制
func findDiagonalOrder(nums [][]int) []int {
	var result []int
	m, n := 0, 0
	flag := false

	// 每行最长长度
	tempMax := 0
	for i := 0; i < len(nums); i++ {
		if len(nums[i]) > tempMax {
			tempMax = len(nums[i])
		}
	}
	max := tempMax + len(nums) - 1

	for i := 0; i < max; i++ {
		// 到了最后一个切片
		if i == len(nums) {
			m = i - 1
			n = i - m
			flag = true
		} else {
			m = i
			n = i - m
		}
		if flag {
			m = len(nums) - 1
			n = i - m
		}

		for m >= 0 {
			// 此处元素为空
			if n < len(nums[m]) {
				result = append(result, nums[m][n])
			}
			m--
			n++
		}
	}

	return result
}

func main()  {
	fmt.Println(findDiagonalOrder([][]int{{6},{8},{6,1,6,16}}))
}