/**
 * @Time : 2020-06-06 23:06
 * @Author : yz
 */

package tree

import (
	"github.com/yinzhenzhen/learning-programming/types"
)

/**
标签：树、深度优先搜索、二叉搜索树、二叉树
难度：中等

题目内容：
98.验证二叉搜索树
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
有效二叉搜索树定义如下：
	节点的左子树只包含小于当前节点的数。
	节点的右子树只包含大于当前节点的数。
	所有左子树和右子树自身必须也是二叉搜索树。

解题思路：见代码注释
        6
     4      8
  1    5  7   9
    2
入栈 6 4 1
出栈 1，栈为 6 4，写入 elems 1
入栈 2，栈为 6 4 2
出栈 2，栈为 6 4，写入 elems 1，2
出栈 4，写入 elems 1，2，4，入栈 5，栈为 6 5
出栈 5，栈为 6，写入 elems 1，2，4，5
出栈 6，写入 elems 1，2，4，5，6，入栈 8
入栈 7，栈为 8，7
出栈 7，栈为 8，写入 elems 1，2，4，5，6，7
出栈 8，写入 elems 1，2，4，5，6，7，8，入栈 9
出栈 9，写入 elems 1，2，4，5，6，7，8，9，结束
*/

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
			// 如果左节点比右节点要大，则不符合规则
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
