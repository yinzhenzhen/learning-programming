/**
 * @Time : 2020-12-01 15:07
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-programming/types"
)

func buildTree(preorder []int, inorder []int) *types.TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &types.TreeNode{}
	if len(preorder) == 1 {
		root.Val = preorder[0]
		return root
	}

	root.Val = preorder[0]

	mid := 0
	for i, v := range inorder {
		if v == root.Val {
			mid = i
			break
		}
	}

	root.Left = buildTree(preorder[1:mid+1], inorder[0:mid])

	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

	return root
}

func main() {
	root := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(root)
}
