/**
 * @Time : 2020-05-22 20:58
 * @Author : yz
 */

package stack

import "github.com/yinzhenzhen/learning-programming/types"

func nextLargerNodes(head *types.ListNode) []int {
	var result []int
	tail := head.Next
	before := 0
	if tail != nil {
		before = tail.Val.(int)
		tail = tail.Next
	}

	for tail != nil {
		val := tail.Val.(int)

		if val > before {

		}

		tail = tail.Next
	}

	return result
}
