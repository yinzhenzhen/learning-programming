/**
 * @Time : 2020-06-06 23:06
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-leetcode/types"
)

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

func main()  {
	elem := []int{1,1}
	root := &types.TreeNode{elem[0], nil, nil}
	root.ConstructTree(elem, 1, 2)
	fmt.Println(isValidBST(root))
}