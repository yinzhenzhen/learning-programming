/**
 * @Time : 2020-05-12 17:00
 * @Author : yz
 */

package stack

import "math"

/**
标签：栈、设计
难度：中等

题目内容：
155.最小栈：
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
实现 MinStack 类:
	MinStack() 初始化堆栈对象。
	void push(int val) 将元素val推入堆栈。
	void pop() 删除堆栈顶部的元素。
	int top() 获取堆栈顶部的元素。
	int getMin() 获取堆栈中的最小元素。

解题思路：见代码注释
*/

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func NewMinStack() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	// 如果新元素比最小栈的栈顶元素要小，则添加x到minStack，否则添加栈顶元素到minStack（最小元素未发生变化）
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(x, top))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
