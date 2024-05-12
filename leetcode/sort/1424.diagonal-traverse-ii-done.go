/**
 * @Time : 2020-06-02 23:10
 * @Author : yz
 */

package sort

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

func findDiagonalOrder2(nums [][]int) []int {
	var result []int
	dict := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			value, exists := dict[i+j]
			if exists {
				value = append(value, nums[i][j])
				dict[i+j] = value
			} else {
				dict[i+j] = []int{nums[i][j]}
			}
		}
	}

	// 每行最长长度
	tempMax := 0
	for i := 0; i < len(nums); i++ {
		if len(nums[i]) > tempMax {
			tempMax = len(nums[i])
		}
	}
	max := tempMax + len(nums) - 1

	for i := 0; i < max; i++ {
		value, _ := dict[i]
		for j := len(value) - 1; j >= 0; j-- {
			result = append(result, value[j])
		}
	}

	return result
}
