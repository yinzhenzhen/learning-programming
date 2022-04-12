/**
 * @Time : 4/12/22 5:40 PM
 * @Author : yz
 */

package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{1, 2, 4, 6, 10}, 8))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
	fmt.Println(twoSum([]int{-1, 0}, -1))
}

func twoSum(numbers []int, target int) []int {

	i, j := 0, len(numbers)-1
	for i < j && i >= 0 && j <= len(numbers)-1 {

		if numbers[i]+numbers[j] == target {
			break
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}

	}
	return []int{i, j}
}
