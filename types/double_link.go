/**
 * @Time : 2020-06-01 20:28
 * @Author : yz
 */

package types

import "fmt"

type DoubleNode struct {
	Val interface{}
	Next *DoubleNode
	Prior *DoubleNode
}

func ConstructDoubleLink(elems []int) (head *DoubleNode) {
	head = &DoubleNode{}
	for _, v := range elems{
		head.TailAppend(v)
	}
	return
}

// 头插法
func (head *DoubleNode) HeadAppend(elem interface{}) {
	node := &DoubleNode{elem, nil, nil}
	node.Next = head.Next
	node.Prior = head
	head.Next = node
}

// 尾插法
func (head *DoubleNode) TailAppend(elem interface{}) {
	node := &DoubleNode{elem, nil, nil}
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = node
	node.Prior = tail
}

func (head *DoubleNode) Remove()  {

}

func (head *DoubleNode) Print()  {
	tail := head.Next
	for tail.Next != nil {
		fmt.Println(tail.Val)
		tail = tail.Next
	}
	fmt.Println(tail.Val)
}
