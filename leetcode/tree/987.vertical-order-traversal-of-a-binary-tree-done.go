/**
 * @Time : 2020-06-08 23:21
 * @Author : yz
 */

package tree

import (
	"github.com/yinzhenzhen/learning-programming/types"
)

type NodeVertical struct {
	node *types.TreeNode
	// 横坐标
	abscissa int
	// 纵坐标
	vertical int
}

func verticalTraversal(root *types.TreeNode) [][]int {
	var result = [][]int{}

	// 每一层的节点, 每层循环会清掉
	var nodes []*types.TreeNode
	nodes = append(nodes, root)

	// 横坐标，和上面对应, 每层循环会清掉
	var abscissa []int
	abscissa = append(abscissa, 0)

	// 纵坐标，每一层节点的垂直值, 和上面对应, 每层循环会清掉
	var vertical []int
	vertical = append(vertical, 0)

	// 所有节点的垂直值
	var nodeVerticals []NodeVertical
	nodeVerticals = append(nodeVerticals, NodeVertical{root, 0, 0})
	min := 0
	max := 0
	flag := true
	for flag {
		var tempNodes []*types.TreeNode
		var tempAbscissa []int
		var tempVertical []int
		// 根据当前这一层的所有节点, 计算下一层节点的垂直值, 并将当前这一层的所有节点写入最后的返回值
		for i := 0; i < len(nodes) && i < len(vertical) && i < len(abscissa); i++ {
			node := nodes[i]
			x := abscissa[i]
			y := vertical[i]
			// 左节点
			leftNode := node.Left
			if leftNode != nil && leftNode.Val != -1 {
				// 如果nodeVerticals已经包含和该节点横纵坐标都相同的节点，保证数值大的节点在后
				tempIndex := len(nodeVerticals)
				for n := 0; n < len(nodeVerticals); n++ {
					nv := nodeVerticals[n]
					if nv.abscissa == x-1 && nv.vertical == y-1 && leftNode.Val < nv.node.Val {
						tempIndex = n
						break
					}
				}
				newNode := NodeVertical{leftNode, x - 1, y - 1}
				temp1 := nodeVerticals[0:tempIndex]
				temp2 := nodeVerticals[tempIndex:]
				// fmt.Println("tempIndex : ", tempIndex)
				// fmt.Println("x-1 : ", x-1)
				// fmt.Println("y-1 : ", y-1)
				// fmt.Println("left node : ", leftNode.Val)
				// fmt.Println("------------------")
				yy := []NodeVertical{}
				nodeVerticals = append(append(append(yy, temp1...), newNode), temp2...)
				tempNodes = append(tempNodes, leftNode)
				tempAbscissa = append(tempAbscissa, x-1)
				tempVertical = append(tempVertical, y-1)
				if x-1 < min {
					min = x - 1
				}
				if x-1 > max {
					max = x - 1
				}
			}
			// 右节点
			rightNode := node.Right
			if rightNode != nil && rightNode.Val != -1 {
				tempIndex := len(nodeVerticals)
				for n := 0; n < len(nodeVerticals); n++ {
					nv := nodeVerticals[n]
					if nv.abscissa == x+1 && nv.vertical == y-1 && rightNode.Val < nv.node.Val {
						tempIndex = n
						break
					}
				}
				newNode := NodeVertical{rightNode, x + 1, y - 1}
				temp1 := nodeVerticals[0:tempIndex]
				temp2 := nodeVerticals[tempIndex:]
				// fmt.Println("tempIndex : ", tempIndex)
				// fmt.Println("x+1 : ", x+1)
				// fmt.Println("y-1 : ", y-1)
				// fmt.Println("right node : ", rightNode.Val)
				// fmt.Println("------------------")
				yy := []NodeVertical{}
				nodeVerticals = append(append(append(yy, temp1...), newNode), temp2...)
				tempNodes = append(tempNodes, rightNode)
				tempAbscissa = append(tempAbscissa, x+1)
				tempVertical = append(tempVertical, y-1)
				if x+1 < min {
					min = x + 1
				}
				if x+1 > max {
					max = x + 1
				}
			}

		}
		nodes = tempNodes
		abscissa = tempAbscissa
		vertical = tempVertical

		i := 0
		for ; i < len(nodes); i++ {
			if nodes[i] != nil {
				break
			}
		}
		if i == len(nodes) {
			flag = false
		}
	}

	// 最后的结果
	m := 0
	for i := min; i <= max; i++ {
		result = append(result, []int{})
		for k := 0; k < len(nodeVerticals); k++ {
			if nodeVerticals[k].abscissa == i {
				result[m] = append(result[m], nodeVerticals[k].node.Val)
			}
		}
		m++
		// fmt.Println("------------------")
	}

	return result
}
