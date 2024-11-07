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
	l := &HeapSqList{nums, len(nums)}
	return l
}

// 堆排序-升序排序
func (l *HeapSqList) HeapSortAsc() {
	// i 为前一半元素的索引，这些元素都有左右子节点，这些元素为关键字
	for i := l.Length/2 - 1; i >= 0; i-- {
		MaxHeapAjust(l, i, l.Length)
	}

	// 输出大顶堆得到升序结果
	for i := l.Length - 1; i > 0; i-- {
		swap(l, 0, i)
		MaxHeapAjust(l, 0, i)
	}
}

// 调整顺序，使s->m的元素成为大顶堆
func MaxHeapAjust(l *HeapSqList, s int, length int) {
	var temp, j int
	temp = l.Data[s]
	// 沿关键字较大的孩子节点向下筛选
	for j = 2*s + 1; j < length; j = j*2 + 1 {
		// s=0 j=1 比较0的子节点1、2的大小，如果2比1大，继续找2的子节点5、6比较他们的大小找到更大的
		if j+1 < length && l.Data[j] < l.Data[j+1] {
			j++
		}
		// 如果0比1和2都大，直接跳出循环
		if temp >= l.Data[j] {
			break
		}
		// 将大的数赋值给0所在的位置
		l.Data[s] = l.Data[j]
		// 数大的索引赋值给s=0
		s = j
	}
	// 将0赋值给大的数所在的位置
	l.Data[s] = temp
}

// 大顶堆
func (l *HeapSqList) MaxHeap() {
	for i := l.Length/2 - 1; i > 0; i-- {
		MaxHeapAjust(l, i, l.Length)
	}
}

// 堆排序-降序排序
func (l *HeapSqList) HeapSortDesc() {
	// i 为前一半元素的索引，这些元素都有左右子节点，这些元素为关键字
	for i := l.Length/2 - 1; i >= 0; i-- {
		MinHeapAjust(l, i, l.Length)
	}

	// 输出大顶堆得到升序结果
	for i := l.Length - 1; i > 0; i-- {
		swap(l, 0, i)
		MinHeapAjust(l, 0, i)
	}
}

// 调整顺序，使s->m的元素成为小顶堆
func MinHeapAjust(l *HeapSqList, s int, length int) {
	var temp, j int
	temp = l.Data[s]
	// 沿关键字较大的孩子节点向下筛选
	for j = 2*s + 1; j < length; j = j*2 + 1 {
		// s=0 j=1 比较0的子节点1、2的大小，如果2比1大，继续找2的子节点5、6比较他们的大小找到更大的
		if j+1 < length && l.Data[j] > l.Data[j+1] {
			j++
		}
		// 如果0比1和2都小，直接跳出循环
		if temp <= l.Data[j] {
			break
		}
		// 将小的数赋值给0所在的位置
		l.Data[s] = l.Data[j]
		// 数小的索引赋值给s
		s = j
	}
	// 将0赋值给小的数所在的位置
	l.Data[s] = temp
}

// 小顶堆
func (l *HeapSqList) MinHeap() {
	for i := l.Length/2 - 1; i > 0; i-- {
		MinHeapAjust(l, i, l.Length)
	}
}

func swap(l *HeapSqList, m int, n int) {
	temp := l.Data[m]
	l.Data[m] = l.Data[n]
	l.Data[n] = temp
}
