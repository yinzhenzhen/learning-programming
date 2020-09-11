/**
 * @Time : 2020-05-13 14:55
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-programming/types"
)


func mergeTrees(t1 *types.TreeNode, t2 *types.TreeNode) *types.TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

func main() {
	elem1 := []int{1,3,2,5}
	tree1 := &types.TreeNode{elem1[0], nil, nil}
	tree1.ConstructTree(elem1, 1, 2)
	elem2 := []int{2,1,3,0,4,0,7}
	tree2 := &types.TreeNode{elem2[0], nil, nil}
	tree2.ConstructTree(elem2, 1, 2)
	fmt.Println(mergeTrees(tree1, tree2))
	fmt.Println()
}