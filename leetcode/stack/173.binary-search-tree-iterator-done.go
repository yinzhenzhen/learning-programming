/**
 * @Time : 2020-05-22 11:17
 * @Author : yz
 */

package stack

import (
	"github.com/yinzhenzhen/learning-programming/types"
)

/**
标签：栈、树、设计、二叉搜索树、二叉树、迭代器
难度：中等

题目内容：
173.二叉搜索树迭代器：
实现一个二叉搜索树迭代器类BSTIterator，表示一个按中序遍历二叉搜索树（BST）的迭代器：
	BSTIterator(TreeNode root) 初始化 BSTIterator 类的一个对象。BST 的根节点 root 会作为构造函数的一部分给出。指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
	boolean hasNext() 如果向指针右侧遍历存在数字，则返回 true；否则返回 false。
	int next()将指针向右移动，然后返回指针处的数字。
注意，指针初始化为一个不存在于 BST 中的数字，所以对 next() 的首次调用将返回 BST 中的最小元素。
你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 的中序遍历中至少存在一个下一个数字。

解题思路：见代码注释
*/

type BSTIterator struct {
	elems []int
	index int
}

func NewBSTIterator(root *types.TreeNode) BSTIterator {
	if root == nil {
		return BSTIterator{index: -1}
	}
	var bst = BSTIterator{index: 0}
	var stack *types.Stack = types.ConstructStack()
	stack.Push(root)
	for stack.Top() != nil {
		top := stack.Top()
		// 如果栈顶元素的左节点不为空，持续加入左节点
		for top.(*types.TreeNode).Left != nil {
			stack.Push(top.(*types.TreeNode).Left)
			top = stack.Top()
		}
		// 栈顶元素没有左节点，弹出该节点
		node := stack.Pop().(*types.TreeNode)
		bst.elems = append(bst.elems, node.Val)
		// 如果弹出的节点有右节点，则将右节点压入栈
		if node.Right != nil {
			stack.Push(node.Right)
		} else {
			// 防止重复取左节点元素，栈顶的左节点设置为空
			if stack.Top() != nil {
				stack.Top().(*types.TreeNode).Left = nil
			}
		}
	}
	return bst
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	if this.index >= len(this.elems) || this.index == -1 {
		return 0
	}
	this.index++
	return this.elems[this.index-1]
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if this.index >= len(this.elems) || this.index == -1 {
		this.elems = []int{}
		this.index = -1
		return false
	}
	return true
}
