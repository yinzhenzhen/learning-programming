/**
 * @Time : 2020-05-28 23:48
 * @Author : yz
 */

package heap

import (
	"fmt"
	"github.com/yinzhenzhen/learning-programming/types"
)

/**
标签：数组、分治、快速选择、排序、堆（优先队列）
难度：中等

题目内容：
215.数组中的第K个最大元素：
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

解题思路：

*/

func findKthLargest(nums []int, k int) int {
	l := types.ConstructHeapSqList(nums)
	l.HeapSortAsc()
	fmt.Println(l.Data)
	return l.Data[l.Length-k]
}
