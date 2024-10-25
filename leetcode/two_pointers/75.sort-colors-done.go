/**
 * @Time : 2020-06-01 10:33
 * @Author : yz
 */

package two_pointers

import (
	"fmt"
	"unsafe"
)

// 任何类型的指针值都可以转换为 Pointer
// Pointer 可以转换为任何类型的指针值
// uintptr 可以转换为 Pointer
// Pointer 可以转换为 uintptr

/**
标签：数组、双指针、排序
难度：中等

题目内容：
75.颜色分类：
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
必须在不使用库内置的 sort 函数的情况下解决这个问题。

解题思路：见代码注释

*/

func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}
	// 指向0的最右端
	p0 := &nums[0]
	// 指向2的最左端
	p2 := &nums[len(nums)-1]
	// 当前指针
	cur := &nums[0]

	for uintptr(unsafe.Pointer(cur)) <= uintptr(unsafe.Pointer(p2)) {
		// 如果当前指向的元素为0, 将该元素和p0指向的元素替换
		if *cur == 0 {
			if *p0 != 0 {
				swap1(cur, p0)
			}
			// cur和p0向右移动
			cur = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(cur)) + uintptr(8)))
			p0 = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p0)) + uintptr(8)))
			// 如果当前指向的元素为2, 将该元素和p2指向的元素替换, 并将p2指针向前移动
		} else if *cur == 2 {
			// 如果当前元素和p2指向的都为2，就不需要替换了
			if *p2 != 2 {
				swap1(cur, p2)
			}
			// p2向左移动
			p2 = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p2)) - uintptr(8)))
		} else {
			// 如果元素为1, cur向右移动
			cur = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(cur)) + uintptr(8)))
		}
	}
	fmt.Println(nums)
}

func swap1(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
