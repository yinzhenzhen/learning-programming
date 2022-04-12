/**
 * @Time : 2020-05-29 09:16
 * @Author : yz
 */

package types

type HeapSqList struct {
	Data   []int
	Length int
}

func ConstructHeapSqList(nums []int) *HeapSqList {
	newNums := []int{}
	newNums = append(newNums, nums...)
	L := &HeapSqList{newNums, len(nums)}
	return L
}

// 堆排序-升序排序
func (L *HeapSqList) HeapSortAsc() {
	// i 为前一半元素的索引，这些元素都有左右子节点，这些元素为关键字
	for i := L.Length/2 - 1; i >= 0; i-- {
		MaxHeapAjust(L, i, L.Length)
	}

	// 输出大顶堆得到升序结果
	for i := L.Length - 1; i > 0; i-- {
		swap(L, 0, i)
		MaxHeapAjust(L, 0, i)
	}
}

func swap(L *HeapSqList, m int, n int) {
	temp := L.Data[m]
	L.Data[m] = L.Data[n]
	L.Data[n] = temp
}

// 调整顺序，使s->m的元素成为大顶堆
func MaxHeapAjust(L *HeapSqList, s int, length int) {
	var temp, j int
	temp = L.Data[s]
	// 沿关键字较大的孩子节点向下筛选
	for j = 2*s + 1; j < length; j = j*2 + 1 {
		// s=0 j=1 比较0的子节点1、2的大小，如果2比1大，继续找2的子节点5、6比较他们的大小找到更大的
		if j+1 < length && L.Data[j] < L.Data[j+1] {
			j++
		}
		// 如果0比1和2都大，直接跳出循环
		if temp >= L.Data[j] {
			break
		}
		// 将大的数赋值给0所在的位置
		L.Data[s] = L.Data[j]
		// 数大的索引赋值给s=0
		s = j
	}
	// 将0赋值给大的数所在的位置
	L.Data[s] = temp
}

// 大顶堆
func (L *HeapSqList) MaxHeap() {
	for i := L.Length/2 - 1; i > 0; i-- {
		MaxHeapAjust(L, i, L.Length)
	}
}

// 堆排序-降序排序
func (L *HeapSqList) HeapSortDesc() {
	// i 为前一半元素的索引，这些元素都有左右子节点，这些元素为关键字
	for i := L.Length/2 - 1; i >= 0; i-- {
		MinHeapAjust(L, i, L.Length)
	}

	// 输出大顶堆得到升序结果
	for i := L.Length - 1; i > 0; i-- {
		swap(L, 0, i)
		MinHeapAjust(L, 0, i)
	}
}

// 小顶堆
func (L *HeapSqList) MinHeap() {
	for i := L.Length/2 - 1; i > 0; i-- {
		MinHeapAjust(L, i, L.Length)
	}
}

// 调整顺序，使s->m的元素成为小顶堆
func MinHeapAjust(L *HeapSqList, s int, length int) {
	var temp, j int
	temp = L.Data[s]
	// 沿关键字较大的孩子节点向下筛选
	for j = 2*s + 1; j < length; j = j*2 + 1 {
		// s=0 j=1 比较0的子节点1、2的大小，如果2比1大，继续找2的子节点5、6比较他们的大小找到更大的
		if j+1 < length && L.Data[j] > L.Data[j+1] {
			j++
		}
		// 如果0比1和2都小，直接跳出循环
		if temp <= L.Data[j] {
			break
		}
		// 将小的数赋值给0所在的位置
		L.Data[s] = L.Data[j]
		// 数小的索引赋值给s
		s = j
	}
	// 将0赋值给小的数所在的位置
	L.Data[s] = temp
}
