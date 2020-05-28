/**
 * @Time : 2020-05-28 23:48
 * @Author : yz
 */

package main

import "fmt"

type SqList struct {
	data []int
	length int
}

func findKthLargest(nums []int, k int) int {
	newNume := []int{0}
	newNume = append(newNume, nums...)
	L := &SqList{newNume, len(nums)}
	HeapSort(L)
	fmt.Println(L.data)
	return L.data[L.length-k+1]
}

// 堆排序
func HeapSort(L *SqList) {
	// i 为前一半元素的索引，这些元素都有左右子节点，这些元素为关键字
	for i := L.length/2; i > 0; i-- {
		HeapAjust(L, i, L.length)
	}

	// 输出大顶堆得到升序结果
	for i := L.length; i > 1; i-- {
		swap(L, 1, i)
		HeapAjust(L, 1, i-1)
	}

}

func swap(L *SqList, m int, n int)  {
	temp := L.data[m]
	L.data[m] = L.data[n]
	L.data[n] = temp
}

// 调整顺序，使s->m的元素成为大顶堆
func HeapAjust(L *SqList, s int, m int)  {
	var temp, j int
	temp = L.data[s]
	// 沿关键字较大的孩子节点向下筛选
	for j = 2*s; j <= m; j*=2 {
		// s=1 j=2 比较1的子节点2、3的大小，如果3比2大，继续找3的子节点6、7比较他们的大小找到更大的
		if j < m && L.data[j] < L.data[j+1] {
			j++
		}
		// 如果1比2和3都大，直接跳出循环
		if temp >= L.data[j] {
			break
		}
		// 将大的数赋值给1所在的位置
		L.data[s] = L.data[j]
		// 数大的索引赋值给s
		s = j
	}
	// 将1赋值给大的数所在的位置
	L.data[s] = temp
}

func main()  {
	fmt.Println(findKthLargest([]int{3,2,3,1,2,4,5,5,6}, 4))
	//fmt.Println(findKthLargest([]int{50,10,90,30,70,40,80,60,20}, 2))
}