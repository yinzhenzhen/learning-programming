/**
 * @Time : 2020-12-01 14:14
 * @Author : yz
 */

package swordfinger

import (
	"github.com/yinzhenzhen/learning-programming/types"
)

/**
从尾到头打印链表
*/
// 递归回溯/辅助栈
func reversePrint(head *types.ListNode) []int {

	if head.Next == nil {
		return nil
	}

	r := reversePrint(head.Next)

	r = append(r, head.Next.Val.(int))

	return r
}
