/**
 * @Time : 2020-06-01 15:03
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-programming/types"
)

// 双向链表
func insertionSortList(head *types.DoubleNode) *types.DoubleNode {
	// 头节点不存值，所以头节点后的第一个节点为首节点，以下为链表为空，or链表只存在一个元素
	if head.Next == nil || head.Next.Next == nil {
		return head
	}

	// 第二个元素
	p := head.Next.Next

	for p != nil  {
		// 当前元素的前一个元素
		q := p.Prior
		// 待比较后进行删除、插入的节点
		s := p
		// 如果当前元素比前一个元素小，继续往前寻找，直到找到比当前元素更小的元素，在该元素之后插入当前元素
		// 跳出循环的条件，要么就是q已经到了头，要么是找到了比当前元素更小的元素
		for q.Val != nil && q.Val.(int) > s.Val.(int) {
			q = q.Prior
		}

		p = p.Next

		// 到了最后一个元素
		if s.Next == nil {
			s.Prior.Next = nil
			s.Prior = nil
		} else {
			// 将s节点从当前位置删除
			s.Prior.Next = s.Next
			s.Next.Prior = s.Prior
		}

		// 将s节点插入寻找到的新的位置
		s.Prior = q
		s.Next = q.Next
		q.Next.Prior = s
		q.Next = s

	}

	return head
}

// 单链表
func insertionSortList2(head *types.ListNode) *types.ListNode {
	// 头节点不存值，所以头节点后的第一个节点为首节点，以下为链表为空，or链表只存在一个元素
	if head.Next == nil || head.Next.Next == nil {
		return head
	}

	// 第二个元素
	p := head.Next.Next

	// 保存s的前一个元素
	m := head.Next

	for p != nil {
		// 第一个元素
		q := head.Next
		// 保存q的前一个元素
		n := head

		// 待比较后进行删除、插入的节点
		s := p

		// 判断m该不该往后移动
		flag := false

		// 如果当前元素比前一个元素小，继续往前寻找，直到找到比当前元素更小的元素，在该元素之后插入当前元素
		// 跳出循环的条件，要么就是q已经到了头，要么是找到了比当前元素更小的元素
		for q.Val != nil && q != s && q.Val.(int) < s.Val.(int) {
			q = q.Next
			n = n.Next
			flag = true
		}

		// s的前一个元素
		//m = s

		if flag {

		}

		// p往下移动
		p = p.Next

		// 到了最后一个元素
		if s.Next == nil {
			m.Next = nil
		} else {
			// 将s节点从当前位置删除
			m.Next = s.Next
		}

		// 将s节点插入寻找到的新的位置
		s.Next = n.Next
		n.Next = s


	}

	return head
}

func main()  {
	//head := types.ConstructDoubleLink([]int{6,5,4,3,2,1})
	//head.Print()
	//fmt.Println()
	//insertionSortList(head)
	//head.Print()

	head := types.ConstructLink([]int{6,5,4,3,2,1})
	head.Print()
	fmt.Println()
	insertionSortList2(head)
	head.Print()
}