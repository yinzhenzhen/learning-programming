/**
 * @Time : 2020-06-10 17:56
 * @Author : yz
 */

package types

type Node struct {
	Val       int
	Neighbors []*Node
}

func ConstructGraph() *Node {
	return &Node{}
}
