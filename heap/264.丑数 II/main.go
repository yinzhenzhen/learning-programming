/**
 * @Time : 2020-05-29 09:13
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-leetcode/types"
)

// 超出时间限制

const (
	two = 2
	three = 3
	five = 5
)

// 丑数，质因数只包含 2, 3, 5 的正整数，求第n个丑数
func nthUglyNumber(n int) int {
	if n <= 6 {
		return n
	}

	// 最终结果的排序
	nums := []int{1}
	// 用于乘以2，3，5
	min := 1
	//
	sort := []int{2,3,5}
	heapSqList := types.ConstructHeapSqList(sort)

	for len(nums) < n {
		// 小顶堆
		heapSqList.HeapSortAsc()
		// 找到小顶堆的最小的元素
		min = heapSqList.Data[1]
		// 删除小的元素
		temp := heapSqList.Data[2:]
		heapSqList.Data = []int{0}
		heapSqList.Data = append(heapSqList.Data, temp...)
		// 新增3个质因数
		heapSqList.Data = append(heapSqList.Data, min*two, min*three, min*five)
		heapSqList.Length = heapSqList.Length+2
		// 去掉重复的元素
		if min == nums[len(nums)-1] {
			continue
		}
		// 将新一轮最小的丑数写入最终的数组
		nums = append(nums, min)

	}

	return nums[n-1]
}

func main()  {
	fmt.Println(nthUglyNumber(95))
}