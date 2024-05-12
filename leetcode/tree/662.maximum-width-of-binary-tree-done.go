/**
 * @Time : 2020-06-07 23:54
 * @Author : yz
 */

package tree

import (
	"github.com/yinzhenzhen/learning-programming/types"
)

//给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。
//这个二叉树与满二叉树（full binary tree）结构相同，但一些节点为空。
//
//每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。

func widthOfBinaryTree(root *types.TreeNode) int {
	if root == nil {
		return 0
	}
	// 宽度最大值
	var maxWidth = 0
	// 每层节点
	var nodes []*types.TreeNode
	nodes = append(nodes, root)
	var flag = true
	for flag {
		var left = -1
		var right = -1
		var leftFlag = false
		var rightFlag = false
		// 计算每层的宽度
		for i, j := 0, len(nodes)-1; !leftFlag || !rightFlag; {
			if left == -1 && nodes[i] != nil && nodes[i].Val != 0 {
				left = i
				leftFlag = true
			}
			if right == -1 && nodes[j] != nil && nodes[j].Val != 0 {
				right = j
				rightFlag = true
			}
			if !leftFlag {
				i++
			}
			if !rightFlag {
				j--
			}
		}
		width := right - left + 1
		// 找到最大宽度值
		if width > maxWidth {
			maxWidth = width
		}
		//fmt.Println("left : ", left)
		//fmt.Println("right : ", right)
		//fmt.Println("width : ", width)

		//for i := 0; i < len(nodes); i++  {
		//	if nodes[i] != nil {
		//		fmt.Println("node : ", nodes[i].Val)
		//	} else {
		//		fmt.Println("node nil")
		//	}
		//}
		//fmt.Println("------------------------")

		// 每层临时节点
		var tempNodes []*types.TreeNode

		// 构造下一层所有节点的集合
		var myFlag = false
		for i := 0; i < len(nodes); i++ {
			node := nodes[i]
			//if node == nil || node.Val == 0 {
			//	tempNodes = append(tempNodes, nil, nil)
			//} else {
			//	tempNodes = append(tempNodes, node.Left, node.Right)
			//}
			if myFlag && node == nil {
				tempNodes = append(tempNodes, nil, nil)
			}
			if node != nil {
				myFlag = true
				tempNodes = append(tempNodes, node.Left, node.Right)
			}
		}
		nodes = tempNodes

		// 如果下一层所有节点都为nil，则结束整个循环
		i := 0
		for ; i < len(nodes); i++ {
			if nodes[i] != nil {
				flag = true
				break
			}
		}
		if i == len(nodes) {
			flag = false
		}
	}

	return maxWidth
}
