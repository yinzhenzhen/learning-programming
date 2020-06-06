/**
 * @Time : 2020-06-06 14:37
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-leetcode/types"
)

func invertTree(root *types.TreeNode) *types.TreeNode {
	if root == nil {
		return root
	}
	swap(root)
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func swap(tree *types.TreeNode)  {
	temp := tree.Left
	tree.Left = tree.Right
	tree.Right = temp
	return
}

func main()  {
	elem := []int{4,2,7,1,3,6,9}
	root := &types.TreeNode{elem[0], nil, nil}
	root.ConstructTree(elem, 1, 2)
	fmt.Println(invertTree(root))
	fmt.Println()
}