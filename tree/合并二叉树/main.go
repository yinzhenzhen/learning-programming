/**
 * @Time : 2020-05-13 14:55
 * @Author : yz
 */

package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
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
	tree1 := &TreeNode{10,
		&TreeNode{5,
			&TreeNode{3, nil, nil},
			&TreeNode{7, nil, nil},
		},
		&TreeNode{15,
			nil,
			&TreeNode{18, nil, nil},
		},
	}

	tree2 := &TreeNode{24,
		&TreeNode{15,
			&TreeNode{12, &TreeNode{9, nil, nil}, nil},
			&TreeNode{21, &TreeNode{18, nil, nil}, nil},
		},
		&TreeNode{33,
			&TreeNode{30, &TreeNode{27, nil, nil}, nil},
			&TreeNode{36, nil, nil},
		},
	}

	fmt.Println(mergeTrees(tree1, tree2))
}