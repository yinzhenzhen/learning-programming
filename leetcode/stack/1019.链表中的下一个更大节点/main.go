/**
 * @Time : 2020-05-22 20:58
 * @Author : yz
 */

package main

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

func main()  {
	head := types.ConstructLink([]int{1, 2, 3})
	head.Print()
}