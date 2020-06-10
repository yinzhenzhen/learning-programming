/**
 * @Time : 2020-06-10 00:27
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-leetcode/types"
)

func constructMaximumBinaryTree(nums []int) *types.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	max := 0
	maxIndex := 0
	// 找最大值
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			maxIndex = i
		}
	}

	root := &types.TreeNode{max, nil, nil}
	root.Left = constructMaximumBinaryTree(nums[0:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])
	return root
}

func main()  {
	root := constructMaximumBinaryTree([]int{3,2,1,6,0,5})
	fmt.Println(root.DFS())
	fmt.Println(root.BFS())
}