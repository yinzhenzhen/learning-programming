package binary_indexed_tree

import (
	"github.com/yinzhenzhen/learning-programming/types"
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
