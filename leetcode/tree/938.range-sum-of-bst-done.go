/**
 * @Time : 2020-05-12 17:05
 * @Author : yz
 */

package tree

import (
	"github.com/yinzhenzhen/learning-programming/types"
)

func rangeSumBST(root *types.TreeNode, L int, R int) int {
	var sum = 0
	inOrderTraversal(root, &sum, L, R)
	return sum
}

func inOrderTraversal(root *types.TreeNode, sum *int, L int, R int) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left, sum, L, R)
	if root.Val >= L && root.Val <= R {
		*sum = *sum + root.Val
	}
	inOrderTraversal(root.Right, sum, L, R)
}
