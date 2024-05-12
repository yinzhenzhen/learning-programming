/**
 * @Time : 2020-05-28 23:48
 * @Author : yz
 */

package heap

import (
	"fmt"
	"github.com/yinzhenzhen/learning-programming/types"
)

func findKthLargest(nums []int, k int) int {
	L := types.ConstructHeapSqList(nums)
	L.HeapSortAsc()
	fmt.Println(L.Data)
	return L.Data[L.Length-k]
}
