/**
 * @Time : 2020-06-12 10:51
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-leetcode/types"
)

func postorderTraversal(root *types.TreeNode) []int {

}

func main()  {
	elem := []int{1,3,2,5}
	root := &types.TreeNode{elem[0], nil, nil}
	root.ConstructTree(elem, 1, 2)
	fmt.Println(postorderTraversal(root))
}