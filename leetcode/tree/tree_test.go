package tree

import (
	"fmt"
	"github.com/yinzhenzhen/learning-programming/types"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	elem := []int{1, 1}
	root := &types.TreeNode{elem[0], nil, nil}
	root.ConstructTree(elem, 1, 2)
	fmt.Println(isValidBST(root))
}

func Test_isSameTree(t *testing.T) {
	elem1 := []int{1, 3, 2, 5}
	tree1 := &types.TreeNode{elem1[0], nil, nil}
	tree1.ConstructTree(elem1, 1, 2)
	elem2 := []int{2, 1, 3, 0, 4, 0, 7}
	tree2 := &types.TreeNode{elem2[0], nil, nil}
	tree2.ConstructTree(elem2, 1, 2)
	fmt.Println(isSameTree(tree1, tree2))
}

func Test_postorderTraversal(t *testing.T) {
	elem := []int{1, 3, 2, 5}
	root := &types.TreeNode{elem[0], nil, nil}
	root.ConstructTree(elem, 1, 2)
	fmt.Println(postorderTraversal(root))
}

func Test_invertTree(t *testing.T) {
	elem := []int{4, 2, 7, 1, 3, 6, 9}
	root := &types.TreeNode{elem[0], nil, nil}
	root.ConstructTree(elem, 1, 2)
	fmt.Println(invertTree(root))
	fmt.Println()
}

func Test_mergeTrees(t *testing.T) {
	elem1 := []int{1, 3, 2, 5}
	tree1 := &types.TreeNode{elem1[0], nil, nil}
	tree1.ConstructTree(elem1, 1, 2)
	elem2 := []int{2, 1, 3, 0, 4, 0, 7}
	tree2 := &types.TreeNode{elem2[0], nil, nil}
	tree2.ConstructTree(elem2, 1, 2)
	fmt.Println(mergeTrees(tree1, tree2))
	fmt.Println()
}

func Test_constructMaximumBinaryTree(t *testing.T) {
	root := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	fmt.Println(root.DFS())
	fmt.Println(root.BFS())
}

func Test_widthOfBinaryTree(t *testing.T) {
	//elem := []int{1,3,2,5,0,0,9,6,0,0,0,0,0,0,7}
	//elem := []int{1,3,0,5,3}
	//elem := []int{2,0,3,0,0,0,1}
	null := -1
	elem := []int{1, 5, 8, 9, 7, 7, 8, 1, 4, 8, 1, 9, 0, 8, 7, 1, 7, 4, 2, 9, 8, 2, 4, null, null, 9, null, null, null, 6, 0, 9, 4, 1, 0, 1, 8, 9, 0, 1, 8, 9, 1,
		0, 9, 6, 2, 5, null, 2, 3, 0, 2, 4, 8, 8, 8, 5, 0, 0, 9, 4, 9, 1, null, 0, 7, 2, 2, 3, null, 6, 1, 0, 8, 9, 9, 9, 4, 8, 4, 3, 4, 4, 0, null, null, 8, 3, 8, null,
		null, 0, null, 0, 4, 9, 1, 2, null, 4, 4, 0, 4, 3, 5, 5, 7, 4, 1, 6, null, 1, 0, null, null, null, 2, 8, 7, 7, null, null, 0, 2, 5, 5, 9, 3, 3, null, 7, 6,
		6, 7, 9, 8, 1, 7, 7, 7, 2, 6, null, 7, null, 4, 6, 4, 6, null, null, 9, 1, null, null, null, 5, 5, 5, 4, 2, 2, 8, 5, 1, 1, 3, 1, 3, 7, null, 2, null, 9, 1, 4, 4, 7, 7,
		null, 1, 5, 6, 2, 7, 3, null, 9, 1, null, 2, 4, 4, 8, null, null, 7, null, 6, null, 7, 4, 3, 5, 8, 4, 8, 5, null, null, 8, null, null, null, 4, 4, null, null,
		null, null, 8, 3, 5, 5, null, null, null, 1, 2, 0, null, null, 9, 3, null, 8, 3, 7, 1, 8, 9, 0, 1, 8, 2, null, 4, null, null, 8, null, null, null, null, 2,
		null, 4, 8, 5, 5, 3, 1, null, null, 6, null, 1, null, null, 6, null, null, null, null, 7, 3, null, null, null, 8, 6, 4, null, 6, 9, 0, 7, 8, null, null,
		0, 6, 7, null, null, 0, 0, 7, 2, 3, 2, null, 0, 2, 3, null, 0, 1, 7, 9, 0, 7, null, null, null, null, 5, 8, 2, 6, 3, 2, 0, 4, null, null, 0, 9, 1, 1, 1, null, 1, 3,
		null, 7, 9, 1, 3, 3, 8, null, null, null, null, 6, null, null, null, null, 9, 8, 1, 3, 8, 3, 0, 6, null, null, 8, 5, 6, 5, 2, 1, null, 5, null, 7, 0, 0,
		null, 9, 3, 9, null, 3, 0, 0, 9, 1, 7, 0, 2, null, 6, 8, 5, null, null, null, null, null, 7, null, 2, 5, null, null, 9, null, null, null, null, null,
		null, null, null, null, null, null, 4, 1, null, 3, 6, 6, 2, 5, 5, 9, null, null, 7, 8, null, null, 2, 7, 3, 7, 2, 5, null, 1, 3, 4, null, null, 8, 3, 6, 9,
		null, 1, null, null, null, null, 9, 7, 5, 2, null, 5, null, 6, 4, 5, null, 1, 2, 0, 6, null, 1, 6, null, null, 5, null, 7, 8, 4, 7, 8, 6, 4, null, 5, 6, 7, 9,
		1, 0, 4, null, null, null, 6, 4, 8, 4, 5, null, 0, 4, 4, 0, 1, 7, 1, null, 1, null, 3, 6, null, null, null, null, 8, null, 5, 0, 7, 5, null, null, 5, 8, null,
		null, 3, null, null, 8, null, 2, 4, null, null, null, null, null, null, null, 9, null, 9, null, 9, null, null, null, null, 7, 1, null, null, 2, null,
		null, 5, 5, 5, 5, 6, 4, null, null, 1, 6, 4, 0, null, 0, 6, 3, 0, null, 5, 5, null, null, null, null, 2, null, 3, 6, null, 3, 0, 5, 0, 1, 0, 3, 4, 9, 9, 2, 7, 3, 8, 6, 9,
		null, 5, 8, null, null, null, null, 9, 8, 0, 7, null, null, 8, 8, 6, 6, 0, 2, 7, 4, 2, 3, 8, 6, 4, null, 8, null, null, null, 2, 0, null, 1, 3, 5, 4, 2, 2, 5, 8,
		8, null, 3, 0, null, 1, 6, 0, null, null, 9, null, 2, null, 6, 8, 2, null, null, 5, null, null, null, 9, 6, 6, 4, 2, 0, null, null, 1, null, 0, null,
		null, null, 6, 6, null, null, null, 4, 7, 9, null, 0, 1, null, null, 9, null, null, null, 4, null, 8, null, null, null, null, null, null, 4, null,
		6, null, 3, null, null, 5, 1, 2, 5, null, 0, 7, 8, null, 7, null, null, 4, null, 4, 4, null, 2, null, 6, null, null, null, 7, null, null, null, null,
		6, 4, null, 6, null, 6, 9, null, null, null, 9, 6, null, 9, null, 3, null, 2, null, 7, 7, null, null, 0, null, 6, 3, null, null, null, null, null,
		null, 1, null, null, null, 6, 9, 7, null, 7, null, 9, 3, 3, null, null, null, null, 4, null, null, 3, null, null, null, 3, 9, null, 0, 3, 1, 9, 6, 7,
		9, 4, 8, null, null, 6, null, 1, 3, 7, null, null, null, 3, null, 2, null, 8, 1, 1, null, null, 6, null, 7, 3, 5, null, 6, 3, 4, null, null, 5, 7, 1, null,
		null, 6, 4, 6, null, null, null, null, 5, 7, 0, 7, 0, null, 5, 8, 5, 5, 4, 5, null, null, null, null, null, null, 1, 7, null, null, 7, null, 9, 9, 6, 4,
		null, null, 3, 2, 1, null, 0, null, 0, 6, null, null, null, 1, 5, null, null, null, 8, null, null, null, null, 3, 4, 8, null, null, 9, 6, 4, null,
		null, null, null, 8, 9, null, 1, null, null, null, 7, null, null, null, null, null, 9, null, null, null, 4, 1, 6, 7, 0, null, null, null, 7, null,
		null, 8, null, null, null, null, null, null, null, 4, null, 9, null, null, null, null, 3, 0, 6, null, 5, null, 9, 9, null, null, 4, 3, 4, null,
		null, null, null, 8, null, 5, null, null, null, null, 5, 2, null, null, null, null, null, null, null, 2, null, null, 2, 1, 8, 5, null, 0, null,
		0, 3, 2, 4, 5, null, null, null, null, null, 7, null, null, 0, null, 0, null, null, null, 0, 3, 9, null, null, null, null, 5, null, null, 0, 5, 0,
		0, null, 9, null, null, null, null, null, null, null, null, 8, null, 9, 3, 5, 9, 0, 5, 9, null, null, 9, 4, null, 0, 2, 0, null, null, 7, null, 7,
		null, 5, 7, 8, 7, null, null, null, 3, 0, 3, null, null, null, null, null, 4, 5, null, null, 2, 3, null, 2, null, null, 7, null, null, 9, null,
		null, 9, 7, 1, null, null, 1, 6, 1, 8, null, null, 5, null, null, 3, 7, 9, 6, null, null, null, null, 1, null, null, null, 3, 7, 3, 2, 3, 3, null, 1,
		null, null, null, 1, null, null, 4, 3, 4, 8, 7, null, 0, 3, 0, null, 1, 1, null, null, null, null, null, 5, null, 6, 0, null, 3, 1, null, 6, null,
		null, 4, 0, 1, null, 6, 1, null, null, 9, 6, 4, 9, 0, 8, 9, 3, 3, 6, null, null, null, null, null, null, null, null, null, null, null, null, 2, null,
		null, null, null, null, 8, 5, 8, 3, 5, 4, null, 6, null, 0, null, null, 6, null, 4, 3, null, null, null, null, null, null, null, null, null, null,
		null, null, null, null, 7, 3, null, null, 1, null, 2, 4, null, null, null, 6, null, null, null, 6, null, 5, null, null, null, null, 1, null, null,
		3, null, 1, null, 7, 1, null, null, 7, 1, 3, 4, 8, null, null, null, null, null, 4, null, null, 4, null, null, null, 7, null, 6, null, null, 1, null,
		null, null, 7, 3, 3, null, null, null, null, 3, 0, null, null, 4, null, null, null, null, null, null, null, null, null, null, 8, null, null, 9,
		null, null, 6, 6, 5, 2, null, 8, 3, 8, null, null, null, null, 6, 7, 0, null, null, null, null, 1, 1, 5, null, 0, 5, null, 5, null, null, null, 1, 2,
		null, 2, 9, 1, null, 2, 4, 1, null, null, null, 1, 8, 4, 4, 5, 2, null, null, 6, 4, 7, 5, 2, 9, null, 4, null, null, null, null, null, 3, null, null, 5,
		9, null, null, null, null, 9, null, 9, null, null, null, 2, null, 1, 9, null, null, null, null, null, 1, 9, 3, null, null, 1, 9, null, 5, 2, 1, 0,
		null, null, 1, 9, 8, 4, 7, null, null, 5, 7, null, null, null, null, 1, 2, 8, null, 6, 0, null, null, null, null, 0, null, null, null, 6, null, 2, 3,
		0, 9, null, null, 1, 4, 6, null, 8, null, null, 5, null, 3, 0, null, 6, null, null, null, null, null, 2, null, null, null, null, null, null, 2, 5, 8,
		6, 9, null, null, null, 8, null, null, 9, 6, null, null, null, null, 3, null, null, null, null, 9, null, null, 2, null, null, null, null, null,
		null, 8, 8, null, null, null, null, null, 9, null, 6, null, 2, 5, null, null, 1, 2, null, 4, null, null, 4, null, null, 3, 5, null, 3, 3, null, null,
		1, null, null, null, null, 4, null, 2, 3, null, 4, 5, 3, null, 7, null, null, null, 7, 6, null, null, 1, 3, null, 4, 9, 8, null, null, 0, null, 3, 4,
		null, 8, null, 1, null, null, 2, 2, null, null, 4, null, null, null, 3, null, null, 2, null, null, null, 4, null, 5, null, null, null, null, 2,
		null, 5, null, null, null, null, null, null, 2, 7, 5, null, 6, null, null, null, null, 2, null, 0, null, 3, null, 1, null, 9, 4, null, 3, null,
		null, null, null, null, null, null, 5, 5, 7, null, null, 1, null, 4, 6, null, null, null, 2, null, 5, 9, 0, 6, 2, null, null, null, null, null,
		null, null, null, null, null, null, null, 5, null, 7, null, 2, 9, null, null, 1, null, null, null, 1, 6, null, 6, null, null, 0, 8, null, 4, null,
		null, null, null, 4, null, null, 0, null, 6, 0, null, null, null, 4, null, null, null, null, null, 0, null, null, null, null, null, null, null,
		null, null, null, null, null, 0, 5, 4, 2, 6, 4, 5, 3, 4, null, null, 5, null, null, null, null, 4, null, null, 3, 6, 2, 0, null, 6, 6, null, null,
		null, null, 0, 6, null, null, null, 3, 9, 4, null, null, null, null, null, 0, null, null, 6, 7, 0, null, 9, 2, null, 3, 3, null, null, 8, null, 3,
		null, null, null, 8, 5, 3, null, 2, 4, null, 9, 6, 9, null, null, null, null, 6, null, 6, null, 5, 3, null, null, null, null, 4, null, null, null,
		9, 0, 9, 7, 1, 1, null, 1, null, 1, 6, null, 5, null, 6, null, null, 1, null, null, null, null, null, null, 5, null, null, null, null, null, 3, null,
		6, 1, null, 0, 2, null, null, 0, null, null, 0, null, null, null, null, null, 3, null, null, 8, null, null, 5, 3, 3, null, null, null, null, null,
		null, null, 3, null, null, 0, 8, 7, null, null, 8, 1, null, null, null, null, null, null, 7, null, null, null, null, null, null, null, null, null,
		null, null, 5, 2, null, 2, 6, null, null, null, null, null, null, null, 1, 5, 0, null, null, 2, null, 7, null, null, 6, null, null, null, null, null,
		null, null, null, null, null, null, null, null, 8, null, null, null, null, 3, null, null, 4, null, null, 2, null, null, null, null, 0, 3,
		null, null, null, null, null, 7, null, 8, null, null, null, null, 8, 5, null, 3, 4, null, null, null, 8, null, null, null, null, null, null,
		null, null, null, 3, 7, null, null, null, 4, 0, 3, null, null, 6, null, null, null, null, null, null, null, null, null, null, null, null, 8,
		null, null, null, null, null, 2, null, null, null, null, null, null, null, null, null, 0, null, null, null, 2, null, null, null, 8, 2, null,
		null, null, null, null, null, null, 8, null, null, null, null, null, null, null, null, null, null, 2, null, null, null, 2, 5, null, null,
		null, null, null, null, null, null, null, null, null, 2, null, null, null, null, null, 8, null, null, null, null, null, null, null,
		null, null, null, 0, 5}
	root := &types.TreeNode{elem[0], nil, nil}
	root.ConstructTree(elem, 1, 2)
	fmt.Println(widthOfBinaryTree(root))
}

func Test_rangeSumBST(t *testing.T) {
	elem := []int{24, 15, 33, 12, 21, 30, 36, 9, 0, 18, 0, 27}
	root := &types.TreeNode{elem[0], nil, nil}
	root.ConstructTree(elem, 1, 2)
	fmt.Println(rangeSumBST(root, 18, 24))
}

func Test_verticalTraversal(t *testing.T) {
	null := -1
	elem := []int{3, 9, 20, null, null, 15, 7}
	//elem := []int{1,2,3,4,5,6,7}
	//elem := []int{0,2,1,3,null,null,null,4,5,null,7,6,null,10,8,11,9}
	root := &types.TreeNode{elem[0], nil, nil}
	root.ConstructTree(elem, 1, 2)
	fmt.Println(verticalTraversal(root))
}

func Test_deepestLeavesSum(t *testing.T) {
	elem := []int{1, 2, 3, 4, 5}
	root := &types.TreeNode{elem[0], nil, nil}
	root.ConstructTree(elem, 1, 2)
	fmt.Println(deepestLeavesSum(root))
}

func Test_maxProduct(t *testing.T) {
	elem := []int{1, 2, 3, 4, 5, 6}
	root := &types.TreeNode{elem[0], nil, nil}
	root.ConstructTree(elem, 1, 2)
	fmt.Println(maxProduct(root))
}
