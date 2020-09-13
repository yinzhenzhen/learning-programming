/**
 * @Time : 2020-09-13 18:41
 * @Author : yz
 */

package main

import "fmt"

type ListNode struct {
	val int
	next *ListNode
}

// 头叉法
func (ln *ListNode) AddNodeByHead(data int)  {

	node := &ListNode{val: data}
	if ln == nil {
		ln = &ListNode{}
		ln.next = node
		return
	}

	p := ln.next
	node.next = p
	ln.next = node
}

func (ln *ListNode) FindKthToTail(nodeIndex int) int {
	p := ln

	for i := 1; i <= nodeIndex && p.next != nil; i++ {
		p = p.next
	}
	return p.val
}

func main()  {
	var num int
	var index int

	for  {
		n, _ := fmt.Scan(&num)
		if n == 0 {
			break
		}

		var ln ListNode
		for i := 0 ; i < num; i++ {

			var temp int
			_, _ = fmt.Scan(&temp)

			ln.AddNodeByHead(temp)

		}

		_, _ = fmt.Scan(&index)

		fmt.Println(ln.FindKthToTail(index))
	}
}