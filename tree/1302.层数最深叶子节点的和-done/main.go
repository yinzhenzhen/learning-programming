/**
 * @Time : 2020-06-10 11:33
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-leetcode/types"
)

func deepestLeavesSum(root *types.TreeNode) int {
	if root == nil {
		return 0
	}

	var nodes []*types.TreeNode
	nodes = append(nodes, root)

	var result int

	for {
		var tempNodes []*types.TreeNode
		var nullNodeNum = 0
		result = 0
		// 每次取出每层内的所有节点，判断该层所有节点的左右节点是否都为空
		for i := 0; i < len(nodes); i++ {
			node := nodes[i]
			if node.Left != nil {
				tempNodes = append(tempNodes, node.Left)
			}
			if node.Right != nil {
				tempNodes = append(tempNodes, node.Right)
			}
			if node.Left == nil && node.Right == nil {
				nullNodeNum++
			}
			result += node.Val
		}
		if nullNodeNum == len(nodes) {
			break
		}
		nodes = tempNodes
	}

	return result
}

func main()  {
	elem := []int{1,2,3,4,5}
	root := &types.TreeNode{elem[0], nil, nil}
	root.ConstructTree(elem, 1, 2)
	fmt.Println(deepestLeavesSum(root))
}