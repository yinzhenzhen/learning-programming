/**
 * @Time : 2020-06-07 00:34
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-leetcode/types"
)

func isSameTree(p *types.TreeNode, q *types.TreeNode) bool {

}

func main()  {
	elem1 := []int{1,3,2,5}
	tree1 := &types.TreeNode{elem1[0], nil, nil}
	tree1.ConstructTree(elem1, 1, 2)
	elem2 := []int{2,1,3,0,4,0,7}
	tree2 := &types.TreeNode{elem2[0], nil, nil}
	tree2.ConstructTree(elem2, 1, 2)
	fmt.Println(isSameTree(tree1, tree2))
}