/**
 * @Time : 2020-06-07 23:28
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-leetcode/types"
)

var sumAll []int

const mod = 1E9 + 7

func maxProduct(root *types.TreeNode) int {
	postOrderTraverse(root)
	result := 0
	for i := 0; i < len(sumAll)-1; i++ {
		result = max(result, sumAll[i] * (sumAll[len(sumAll)-1] - sumAll[i]))
	}
	sumAll = []int{}
	return result % mod
}

func max(a, b int) int {
	if a > b {
		return a
	}
	if b > a {
		return b
	}
	return a
}

func postOrderTraverse(root *types.TreeNode) int {
	if root == nil {
		return 0
	}
	sum := root.Val + postOrderTraverse(root.Left) + postOrderTraverse(root.Right)
	sumAll = append(sumAll, sum)
	return sum
}

func main()  {
	elem := []int{1,2,3,4,5,6}
	root := &types.TreeNode{elem[0], nil, nil}
	root.ConstructTree(elem, 1, 2)
	fmt.Println(maxProduct(root))
}