/**
 * @Time : 2020-05-22 09:53
 * @Author : yz
 */

package types

import "fmt"

// 二叉树
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func (tree *TreeNode) ConstructTree(elems []int, leftIndex int, rightIndex int)  {
	if leftIndex < len(elems) {
		left := &TreeNode{}
		left.Val = elems[leftIndex]
		tree.Left = left
		tree.Left.ConstructTree(elems, 2*leftIndex+1, 2*leftIndex+2)
	}
	if rightIndex < len(elems) {
		right := &TreeNode{}
		right.Val = elems[rightIndex]
		tree.Right = right
		tree.Right.ConstructTree(elems, 2*rightIndex+1, 2*rightIndex+2)
	}
	return
}

// 前序遍历
func (tree *TreeNode) PreOrderTraverse()  {
	if tree == nil {
		return
	}
	fmt.Printf("%d\n", tree.Val)
	tree.Left.PreOrderTraverse()
	tree.Right.PreOrderTraverse()
}

// 中序遍历
func (tree *TreeNode) InOrderTraverse()  {
	if tree == nil {
		return
	}
	tree.Left.InOrderTraverse()
	fmt.Printf("%d\n", tree.Val)
	tree.Right.InOrderTraverse()
}

// 后序遍历
func (tree *TreeNode) PostOrderTraverse()  {
	if tree == nil {
		return
	}
	tree.Left.PostOrderTraverse()
	tree.Right.PostOrderTraverse()
	fmt.Printf("%d\n", tree.Val)
}
