/**
 * @Time : 2020-05-22 20:59
 * @Author : yz
 */

package types

import "fmt"

type ListNode struct {
	Val interface{}
	Next *ListNode
}

func ConstructLink(elems []int) (head *ListNode) {
	head = &ListNode{}
	for _, v := range elems{
		head.TailAppend(v)
	}
	return
}

// 头插法
func (head *ListNode) HeadAppend(elem interface{}) {
	node := &ListNode{elem, nil}
	node.Next = head.Next
	head.Next = node
}

// 尾插法
func (head *ListNode) TailAppend(elem interface{}) {
	node := &ListNode{elem, nil}
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = node
}

func (head *ListNode) Remove(elem interface{})  {

}

func (head *ListNode) Print()  {
	tail := head.Next
	for tail.Next != nil {
		fmt.Println(tail.Val)
		tail = tail.Next
	}
	fmt.Println(tail.Val)
}
