/**
 * @Time : 2020-12-01 14:14
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-programming/types"
)

// 递归回溯/辅助栈
func reversePrint(head *types.ListNode) []int {

	if head.Next == nil {
		return nil
	}

	r := reversePrint(head.Next)

	r = append(r, head.Next.Val.(int))

	return r
}

func main() {
	head := types.ConstructLink([]int{1, 3, 2})
	fmt.Println(reversePrint(head))
}
