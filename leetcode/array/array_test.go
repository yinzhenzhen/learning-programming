package array

import (
	"fmt"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	nums := []int{2, 3, 5}
	target := 8
	result := twoSum(nums, target)
	fmt.Println(result)

	nums = []int{2, 7, 11, 15}
	target = 9
	result = twoSum(nums, target)
	fmt.Println(result)
}
