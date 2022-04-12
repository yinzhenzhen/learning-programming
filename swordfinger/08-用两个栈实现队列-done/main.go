/**
 * @Time : 2020-12-01 17:04
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-programming/types"
)

type CQueue struct {
	// 主要存放元素的栈
	s1 *types.Stack
	// 辅助栈
	s2 *types.Stack
}

func Constructor() CQueue {
	return CQueue{
		s1: &types.Stack{},
		s2: &types.Stack{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.s1.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.s1.Elem) == 0 {
		return -1
	}

	// 将所有元素放入辅助栈
	for len(this.s1.Elem) > 1 {
		this.s2.Push(this.s1.Pop())
	}

	result := this.s1.Pop()

	// 将辅助栈的元素重新放入主栈
	for len(this.s2.Elem) > 0 {
		this.s1.Push(this.s2.Pop())
	}

	return result.(int)
}

func main() {

	obj := Constructor()
	obj.AppendTail(1)
	obj.AppendTail(2)
	p1 := obj.DeleteHead()
	p2 := obj.DeleteHead()
	fmt.Println(p1, p2)
}
