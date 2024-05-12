package swordfinger

import (
	"fmt"
	"github.com/yinzhenzhen/learning-programming/types"
	"testing"
)

func Test_findRepeatNumber(t *testing.T) {
	fmt.Println(findRepeatNumber([]int{1, 2, 3, 4, 5, 1}))
}

func Test_findNumberIn2DArray(t *testing.T) {
	fmt.Println(findNumberIn2DArray([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30}}, 20))
}

func Test_replaceSpace(t *testing.T) {
	fmt.Println(replaceSpace("We are happy."))
}

func Test_reversePrint(t *testing.T) {
	head := types.ConstructLink([]int{1, 3, 2})
	fmt.Println(reversePrint(head))
}

func Test_buildTree(t *testing.T) {
	root := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(root)
}

func Test_CQueue(t *testing.T) {
	obj := NewCQueue()
	obj.AppendTail(1)
	obj.AppendTail(2)
	p1 := obj.DeleteHead()
	p2 := obj.DeleteHead()
	fmt.Println(p1, p2)
}
