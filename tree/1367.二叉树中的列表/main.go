/**
 * @Time : 2020-06-08 11:05
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-leetcode/types"
)

func isSubPath(head *types.ListNode, root *types.TreeNode) bool {

}

func main()  {
	elem := []int{2,0,3,0,0,0,1}
	root := &types.TreeNode{elem[0], nil, nil}
	root.ConstructTree(elem, 1, 2)
	list := []int{4,2,8}
	link := types.ConstructLink(list)
	fmt.Println(isSubPath(link, root))
}