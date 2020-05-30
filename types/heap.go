/**
 * @Time : 2020-05-29 09:16
 * @Author : yz
 */

package types

type HeapSqList struct {
	Data []int
	Length int
}

func ConstructHeapSqList(nums []int) *HeapSqList {
	// 第一个位置不存储元素
	newNums := []int{0}
	newNums = append(newNums, nums...)
	L := &HeapSqList{newNums, len(nums)}
	return L
}

// 堆排序-升序排序
func (L *HeapSqList) HeapSortAsc() {
	// i 为前一半元素的索引，这些元素都有左右子节点，这些元素为关键字
	for i := L.Length/2; i > 0; i-- {
		MaxHeapAjust(L, i, L.Length)
	}

	// 输出大顶堆得到升序结果
	for i := L.Length; i > 1; i-- {
		swap(L, 1, i)
		MaxHeapAjust(L, 1, i-1)
	}
}

func swap(L *HeapSqList, m int, n int)  {
	temp := L.Data[m]
	L.Data[m] = L.Data[n]
	L.Data[n] = temp
}

// 调整顺序，使s->m的元素成为大顶堆
func MaxHeapAjust(L *HeapSqList, s int, m int)  {
	var temp, j int
	temp = L.Data[s]
	// 沿关键字较大的孩子节点向下筛选
	for j = 2*s; j <= m; j*=2 {
		// s=1 j=2 比较1的子节点2、3的大小，如果3比2大，继续找3的子节点6、7比较他们的大小找到更大的
		if j < m && L.Data[j] < L.Data[j+1] {
			j++
		}
		// 如果1比2和3都大，直接跳出循环
		if temp >= L.Data[j] {
			break
		}
		// 将大的数赋值给1所在的位置
		L.Data[s] = L.Data[j]
		// 数大的索引赋值给s
		s = j
	}
	// 将1赋值给大的数所在的位置
	L.Data[s] = temp
}

// 大顶堆
func (L *HeapSqList) MaxHeap()  {
	for i := L.Length/2; i > 0; i-- {
		MaxHeapAjust(L, i, L.Length)
	}
}

// 堆排序-降序排序
func (L *HeapSqList) HeapSortDesc() {
	// i 为前一半元素的索引，这些元素都有左右子节点，这些元素为关键字
	for i := L.Length/2; i > 0; i-- {
		MinHeapAjust(L, i, L.Length)
	}

	// 输出大顶堆得到升序结果
	for i := L.Length; i > 1; i-- {
		swap(L, 1, i)
		MinHeapAjust(L, 1, i-1)
	}
}

// 小顶堆
func (L *HeapSqList) MinHeap()  {
	for i := L.Length/2; i > 0; i-- {
		MinHeapAjust(L, i, L.Length)
	}
}


// 调整顺序，使s->m的元素成为小顶堆
func MinHeapAjust(L *HeapSqList, s int, m int)  {
	var temp, j int
	temp = L.Data[s]
	// 沿关键字较大的孩子节点向下筛选
	for j = 2*s; j <= m; j*=2 {
		// s=1 j=2 比较1的子节点2、3的大小，如果3比2小，继续找3的子节点6、7比较他们的大小找到更小的
		if j < m && L.Data[j] > L.Data[j+1] {
			j++
		}
		// 如果1比2和3都小，直接跳出循环
		if temp <= L.Data[j] {
			break
		}
		// 将小的数赋值给1所在的位置
		L.Data[s] = L.Data[j]
		// 数大的索引赋值给s
		s = j
	}
	// 将1赋值给小的数所在的位置
	L.Data[s] = temp
}
