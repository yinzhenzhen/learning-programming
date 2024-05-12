package binary_indexed_tree

import (
	"fmt"
	"sort"

	"github.com/yinzhenzhen/learning-programming/types"
)

func countSmaller(nums []int) []int {
	// copy 一份原数组至所有数字 allNums 数组中
	allNums, res := make([]int, len(nums)), []int{}
	copy(allNums, nums)
	// 将 allNums 离散化
	sort.Ints(allNums)
	k := 1
	kth := map[int]int{allNums[0]: k}
	for i := 1; i < len(allNums); i++ {
		if allNums[i] != allNums[i-1] {
			k++
			kth[allNums[i]] = k
		}
	}
	fmt.Println(kth)
	// 树状数组 Query
	bit := types.BinaryIndexedTree{}
	bit.Init(k)
	for i := len(nums) - 1; i >= 0; i-- {
		res = append(res, bit.Query(kth[nums[i]]-1))
		bit.Add(kth[nums[i]], 1)
	}
	fmt.Println(bit)
	fmt.Println(res)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

func countSmaller2(nums []int) []int {
	res := make([]int, len(nums))

	res[len(nums)-1] = 0

	for i := len(nums) - 2; i >= 0; i-- {

		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[i] > nums[j] {
				res[i]++
			}
		}
	}

	return res
}
