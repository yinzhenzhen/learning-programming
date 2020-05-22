/**
 * @Time : 2020-05-12 17:05
 * @Author : yz
 */

package main

import "fmt"

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

var res []int

//func rangeSumBST(root *TreeNode, L int, R int) int {
//	var r int
//	inOrderTraversal(root)
//	//fmt.Println()
//	//fmt.Println(res)
//	for _, v := range res {
//		if v >= L && v <= R{
//			r += v
//		}
//	}
//	return r
//}

func rangeSumBST(root *TreeNode, L int, R int) int {
	var r int
	var flag bool = false
	inOrderTraversal(root)
	//fmt.Println()
	//fmt.Println(res)
	for _, v := range res {
		if v == L {
			r += v
			flag = true
			continue
		}
		if flag {
			r += v
			if v == R {
				break
			}
		}
	}
	return r
}

func inOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left)
	fmt.Printf("%d -> ", root.Val)
	res = append(res, root.Val)
	inOrderTraversal(root.Right)
}

func main()  {
	//tree := &TreeNode{10,
	//	&TreeNode{5,
	//		&TreeNode{3, nil, nil},
	//		&TreeNode{7, nil, nil},
	//	},
	//	&TreeNode{15,
	//		nil,
	//		&TreeNode{18, nil, nil},
	//	},
	//}

	tree := &TreeNode{24,
		&TreeNode{15,
			&TreeNode{12, &TreeNode{9, nil, nil}, nil},
			&TreeNode{21, &TreeNode{18, nil, nil}, nil},
		},
		&TreeNode{33,
			&TreeNode{30, &TreeNode{27, nil, nil}, nil},
			&TreeNode{36, nil, nil},
		},
	}

	fmt.Println(rangeSumBST(tree, 18, 24))
}