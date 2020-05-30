/**
 * @Time : 2020-05-28 23:48
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-leetcode/types"
)



func findKthLargest(nums []int, k int) int {
	L := types.ConstructHeapSqList(nums)
	L.HeapSortAsc()
	fmt.Println(L.Data)
	return L.Data[L.Length-k+1]
}

func main()  {
	fmt.Println(findKthLargest([]int{3,2,3,1,2,4,5,5,6}, 4))
	//fmt.Println(findKthLargest([]int{50,10,90,30,70,40,80,60,20}, 2))
}