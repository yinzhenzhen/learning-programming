/**
 * @Time : 2020-11-30 15:51
 * @Author : yz
 */

package swordfinger

/*
*
数组中重复的数字
*/
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
