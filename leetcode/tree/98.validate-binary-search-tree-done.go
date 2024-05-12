/**
 * @Time : 2020-06-06 23:06
 * @Author : yz
 */

package tree

import (
	"github.com/yinzhenzhen/learning-programming/types"
)

//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//假设一个二叉搜索树具有如下特征：
//	节点的左子树只包含小于当前节点的数。
//	节点的右子树只包含大于当前节点的数。
//	所有左子树和右子树自身必须也是二叉搜索树。

func isValidBST(root *types.TreeNode) bool {
	if root == nil {
		return true
	}
	var stack types.Stack
	var elems []int
	stack.Push(root)
	for !stack.IsEmpty() {
		// 栈顶元素
		top := stack.Top().(*types.TreeNode)
		// 如果栈顶元素还有左节点，持续加入左节点，栈顶元素的左节点要么为空，要么是已经被置为空
		for top.Left != nil {
			stack.Push(top.Left)
			top = stack.Top().(*types.TreeNode)
		}
		node := stack.Pop().(*types.TreeNode)
		elems = append(elems, node.Val)
		// 比较节点大小
		if len(elems) > 1 && elems[len(elems)-2] >= elems[len(elems)-1] {
			return false
		}
		// 如果弹出的节点有右节点，则将右节点压入栈
		if node.Right != nil {
			stack.Push(top.Right)
		} else {
			// 防止重复取左节点元素，栈顶的左节点设置为空
			if stack.Top() != nil {
				stack.Top().(*types.TreeNode).Left = nil
			}
		}
	}

	return true
}
