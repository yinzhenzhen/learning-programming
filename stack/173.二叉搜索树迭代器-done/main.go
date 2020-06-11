/**
 * @Time : 2020-05-22 11:17
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-leetcode/types"
)

//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}

type BSTIterator struct {
	elems []int
	index int
}

func Constructor(root *types.TreeNode) BSTIterator {
	if root == nil {
		return BSTIterator{index:-1}
	}
	var bst = BSTIterator{index:0}
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

func main()  {
	null := -1
	elem := []int{7,3,15,null,null,9,20}
	root := &types.TreeNode{elem[0], nil, nil}
	root.ConstructTree(elem, 1, 2)
	bst := Constructor(root)
	fmt.Println(bst.elems)
	fmt.Println(bst.Next())
	fmt.Println(bst.Next())
	fmt.Println(bst.HasNext())
	fmt.Println(bst.Next())
	fmt.Println(bst.Next())
	fmt.Println(bst.Next())
	fmt.Println(bst.Next())
	fmt.Println(bst.Next())
	fmt.Println(bst.HasNext())
	fmt.Println(bst.Next())
}