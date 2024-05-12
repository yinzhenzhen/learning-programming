package queue

import (
	"fmt"
	"testing"
)

func Test_leastInterval(t *testing.T) {
	//fmt.Println(leastInterval([]byte{97,97,97,97,97,97,98,99,100,101,102,103}, 2))
	//fmt.Println(leastInterval([]byte{97,98,99,100,97,98,103}, 3))
	fmt.Println(leastInterval([]byte{97, 97, 97, 97, 98, 98, 98, 98, 99, 99, 99, 99, 100, 100, 100, 100, 101, 102}, 4))
}

func Test_MyCircularQueue(t *testing.T) {
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
