/**
 * @Time : 2020-05-31 19:56
 * @Author : yz
 */

package main

import "fmt"

// 队列长度
var maxSize int = 0

type MyCircularQueue struct {
	data []int
	// 队首索引
	front int
	// 队尾索引
	rear int
}


/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	cq := MyCircularQueue{}
	maxSize = k+1
	cq.front = 0
	cq.rear = 0
	return cq
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.rear < this.front && this.front != 0 {
		this.data[this.rear] = value
	} else {
		this.data = append(this.data, value)
	}
	this.rear = (this.rear+1) % (maxSize)
	return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.data[this.front] = 0
	this.front = (this.front+1) % maxSize
	return true
}


/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.front]
}


/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	if this.rear == 0  {
		return this.data[maxSize-1]
	}
	return this.data[this.rear-1]
}


/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if this.rear == this.front {
		return true
	}
	return false
}


/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if (this.rear+1) % maxSize == this.front {
		return true
	}
	return false
}

func main()  {
	q := Constructor(3)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	fmt.Println(q.data)
	fmt.Println(q.Rear())
	fmt.Println(q.IsFull())
	q.DeQueue()
	fmt.Println(q.data)
	q.EnQueue(4)
	fmt.Println(q.Rear())
	fmt.Println(q.data)
	q.DeQueue()
	fmt.Println(q.data)
	q.EnQueue(5)
	fmt.Println(q.data)
	//fmt.Println(q.data)
	//fmt.Println(q.front)
	//fmt.Println(q.rear)
	//fmt.Println(q.Front())

	//q.EnQueue(6)
	//fmt.Println(q.IsEmpty())
	//fmt.Println(q.IsFull())
}