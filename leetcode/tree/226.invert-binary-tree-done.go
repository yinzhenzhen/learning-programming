/**
 * @Time : 2020-06-06 14:37
 * @Author : yz
 */

package tree

import (
	"github.com/yinzhenzhen/learning-programming/types"
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

func swap(tree *types.TreeNode) {
	temp := tree.Left
	tree.Left = tree.Right
	tree.Right = temp
	return
}
