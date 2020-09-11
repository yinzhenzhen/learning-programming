/**
 * @Time : 2020-05-12 17:05
 * @Author : yz
 */

package main

import (
	"fmt"
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

func main()  {
	elem := []int{24,15,33,12,21,30,36,9,0,18,0,27}
	root := &types.TreeNode{elem[0], nil, nil}
	root.ConstructTree(elem, 1, 2)
	fmt.Println(rangeSumBST(root, 18, 24))
}