/**
 * @Time : 2020-06-06 23:06
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-leetcode/types"
	"math"
)

func isValidBST(root *types.TreeNode) bool {
	var stack types.Stack
	var min = math.MinInt64
	for root != nil || !stack.IsEmpty()  {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		root = stack.Pop().(*types.TreeNode)
		if root.Val < min {
			return false
		}
		min = root.Val
		root = root.Right
	}

	return true
}

func main()  {
	elem := []int{10,5,15,0,0,6,20}
	root := &types.TreeNode{elem[0], nil, nil}
	root.ConstructTree(elem, 1, 2)
	fmt.Println(isValidBST(root))
}