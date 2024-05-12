/**
 * @Time : 2022/6/29 9:57 上午
 * @Author : yz
 */

package two_pointers

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	start, end, length := 0, 0, len(nums)
	result := make([][]int, 0)

	for index := 1; index < length-1; index++ {
		start, end = 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			sum := nums[start] + nums[index] + nums[end]
			if sum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if sum > 0 {
				end--
			} else {
				start++
			}
		}
	}

	return result
}
