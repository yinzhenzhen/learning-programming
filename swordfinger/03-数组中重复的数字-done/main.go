/**
 * @Time : 2020-11-30 15:51
 * @Author : yz
 */

package main

import "fmt"

func findRepeatNumber(nums []int) int {
	repeat := make(map[int]bool)

	for _, v := range nums {
		if _, ok := repeat[v]; ok {
			return v
		}
		repeat[v] = true
	}

	return 0
}

func main() {
	fmt.Println(findRepeatNumber([]int{1, 2, 3, 4, 5, 1}))
}
