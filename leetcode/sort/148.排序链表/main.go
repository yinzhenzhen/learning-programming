/**
 * @Time : 2020-06-02 00:27
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-programming/types"
)

// 使用归并排序
func sortList(head *types.ListNode) *types.ListNode {

}

func main()  {
	head := types.ConstructLink([]int{6,5,4,3,2,1})
	head.Print()
	fmt.Println()
	sortList(head)
	head.Print()
}