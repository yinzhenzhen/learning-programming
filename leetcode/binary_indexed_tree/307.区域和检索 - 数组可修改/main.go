package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-leetcode/types"
)

type NumArray struct {
	bit  types.BinaryIndexedTree
	data []int
}

// Constructor define
func Constructor(nums []int) NumArray {
	bit := types.BinaryIndexedTree{}
	bit.InitWithNums(nums)
	return NumArray{bit: bit, data: nums}
}

// Update define
func (this *NumArray) Update(i int, val int) {
	this.bit.Add(i+1, val-this.data[i])
	this.data[i] = val
}

// SumRange define
func (this *NumArray) SumRange(i int, j int) int {
	return this.bit.Query(j+1) - this.bit.Query(i)
}

func main() {
	nums := []int{1, 3, 5}
	obj := Constructor(nums)
	fmt.Println(obj)
	fmt.Println(obj.SumRange(0, 2))
	obj.Update(1, 2)
	fmt.Println(obj.SumRange(0, 2))
}
