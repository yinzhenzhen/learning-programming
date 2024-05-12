package binary_indexed_tree

import (
	"fmt"
	"testing"
)

func Test_getSkyline(t *testing.T) {
	fmt.Println(getSkyline([][]int{}))
}

func Test_SumRange(t *testing.T) {
	nums := []int{1, 3, 5}
	obj := Constructor(nums)
	fmt.Println(obj)
	fmt.Println(obj.SumRange(0, 2))
	obj.Update(1, 2)
	fmt.Println(obj.SumRange(0, 2))
}

func Test_countSmaller2(t *testing.T) {
	fmt.Println(countSmaller2([]int{5, 2, 6, 1}))
}
