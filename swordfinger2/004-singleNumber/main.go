/**
 * @Time : 4/12/22 3:36 PM
 * @Author : yz
 */

package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 100}))
}

/*
*
只出现一次的数字
*/
func singleNumber(nums []int) int {

	m := make(map[int]int)
	result := 0

	for i := 0; i < len(nums); i++ {
		_, ok := m[nums[i]]
		if !ok {
			m[nums[i]] = 1
		} else {
			m[nums[i]]++
		}
	}

	for k, v := range m {
		if v == 1 {
			result = k
			break
		}
	}

	return result
}
