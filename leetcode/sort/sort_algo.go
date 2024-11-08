package sort

import (
	"math"
	"sort"
)

/* 选择排序 */
func selectionSort(nums []int) {
	n := len(nums)
	// 外循环：未排序区间为 [i, n-1]
	for i := 0; i < n-1; i++ {
		// 内循环：找到未排序区间内的最小元素
		k := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[k] {
				// 记录最小元素的索引
				k = j
			}
		}
		// 将该最小元素与未排序区间的首个元素交换
		nums[i], nums[k] = nums[k], nums[i]
	}
}

/* 冒泡排序 */
func bubbleSort(nums []int) {
	// 外循环：未排序区间为 [0, i]
	for i := len(nums) - 1; i > 0; i-- {
		// 内循环：将未排序区间 [0, i] 中的最大元素交换至该区间的最右端
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				// 交换 nums[j] 与 nums[j + 1]
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

/* 冒泡排序（标志优化）*/
func bubbleSortWithFlag(nums []int) {
	// 外循环：未排序区间为 [0, i]
	for i := len(nums) - 1; i > 0; i-- {
		flag := false // 初始化标志位
		// 内循环：将未排序区间 [0, i] 中的最大元素交换至该区间的最右端
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				// 交换 nums[j] 与 nums[j + 1]
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true // 记录交换元素
			}
		}
		if flag == false { // 此轮“冒泡”未交换任何元素，直接跳出
			break
		}
	}
}

/* 插入排序 */
func insertionSort(nums []int) {
	// 外循环：已排序区间为 [0, i-1]
	for i := 1; i < len(nums); i++ {
		base := nums[i]
		j := i - 1
		// 内循环：将 base 插入到已排序区间 [0, i-1] 中的正确位置
		for j >= 0 && nums[j] > base {
			nums[j+1] = nums[j] // 将 nums[j] 向右移动一位
			j--
		}
		nums[j+1] = base // 将 base 赋值到正确位置
	}
}

type quickSort struct {
}

/* 快速排序 */
func (q *quickSort) quickSort(nums []int, left, right int) {
	// 子数组长度为 1 时终止递归
	if left >= right {
		return
	}
	// 哨兵划分
	pivot := q.partition(nums, left, right)
	// 递归左子数组、右子数组
	q.quickSort(nums, left, pivot-1)
	q.quickSort(nums, pivot+1, right)
}

/* 哨兵划分 */
func (q *quickSort) partition(nums []int, left, right int) int {
	// 以 nums[left] 为基准数
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j-- // 从右向左找首个小于基准数的元素
		}
		for i < j && nums[i] <= nums[left] {
			i++ // 从左向右找首个大于基准数的元素
		}
		// 元素交换
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 将基准数交换至两子数组的分界线
	nums[i], nums[left] = nums[left], nums[i]
	return i // 返回基准数的索引
}

/* 归并排序 */
func mergeSort(nums []int, left, right int) {
	// 终止条件
	if left >= right {
		return
	}
	// 划分阶段
	mid := left + (right-left)/2
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)
	// 合并阶段
	merge(nums, left, mid, right)
}

/* 合并左子数组和右子数组 */
func merge(nums []int, left, mid, right int) {
	// 左子数组区间为 [left, mid], 右子数组区间为 [mid+1, right]
	// 创建一个临时数组 tmp ，用于存放合并后的结果
	tmp := make([]int, right-left+1)
	// 初始化左子数组和右子数组的起始索引
	i, j, k := left, mid+1, 0
	// 当左右子数组都还有元素时，进行比较并将较小的元素复制到临时数组中
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}
	// 将左子数组和右子数组的剩余元素复制到临时数组中
	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}
	for j <= right {
		tmp[k] = nums[j]
		j++
		k++
	}
	// 将临时数组 tmp 中的元素复制回原数组 nums 的对应区间
	for k := 0; k < len(tmp); k++ {
		nums[left+k] = tmp[k]
	}
}

/* 堆排序 */
func heapSort(nums *[]int) {
	// 建堆操作：堆化除叶节点以外的其他所有节点
	for i := len(*nums)/2 - 1; i >= 0; i-- {
		siftDown(nums, len(*nums), i)
	}
	// 从堆中提取最大元素，循环 n-1 轮
	for i := len(*nums) - 1; i > 0; i-- {
		// 交换根节点与最右叶节点（交换首元素与尾元素）
		(*nums)[0], (*nums)[i] = (*nums)[i], (*nums)[0]
		// 以根节点为起点，从顶至底进行堆化
		siftDown(nums, i, 0)
	}
}

/* 堆的长度为 n ，从节点 i 开始，从顶至底堆化 */
func siftDown(nums *[]int, n, i int) {
	for true {
		// 判断节点 i, l, r 中值最大的节点，记为 ma
		l := 2*i + 1
		r := 2*i + 2
		ma := i
		if l < n && (*nums)[l] > (*nums)[ma] {
			ma = l
		}
		if r < n && (*nums)[r] > (*nums)[ma] {
			ma = r
		}
		// 若节点 i 最大或索引 l, r 越界，则无须继续堆化，跳出
		if ma == i {
			break
		}
		// 交换两节点
		(*nums)[i], (*nums)[ma] = (*nums)[ma], (*nums)[i]
		// 循环向下堆化
		i = ma
	}
}

/* 桶排序 */
func bucketSort(nums []float64) {
	// 初始化 k = n/2 个桶，预期向每个桶分配 2 个元素
	k := len(nums) / 2
	buckets := make([][]float64, k)
	for i := 0; i < k; i++ {
		buckets[i] = make([]float64, 0)
	}
	// 1. 将数组元素分配到各个桶中
	for _, num := range nums {
		// 输入数据范围为 [0, 1)，使用 num * k 映射到索引范围 [0, k-1]
		i := int(num * float64(k))
		// 将 num 添加进桶 i
		buckets[i] = append(buckets[i], num)
	}
	// 2. 对各个桶执行排序
	for i := 0; i < k; i++ {
		// 使用内置切片排序函数，也可以替换成其他排序算法
		sort.Float64s(buckets[i])
	}
	// 3. 遍历桶合并结果
	i := 0
	for _, bucket := range buckets {
		for _, num := range bucket {
			nums[i] = num
			i++
		}
	}
}

/* 计数排序 */
// 简单实现，无法用于排序对象
func countingSortNaive(nums []int) {
	// 1. 统计数组最大元素 m
	m := 0
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	// 2. 统计各数字的出现次数
	// counter[num] 代表 num 的出现次数
	counter := make([]int, m+1)
	for _, num := range nums {
		counter[num]++
	}
	// 3. 遍历 counter ，将各元素填入原数组 nums
	for i, num := 0, 0; num < m+1; num++ {
		for j := 0; j < counter[num]; j++ {
			nums[i] = num
			i++
		}
	}
}

/* 计数排序 */
// 完整实现，可排序对象，并且是稳定排序
func countingSort(nums []int) {
	// 1. 统计数组最大元素 m
	m := 0
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	// 2. 统计各数字的出现次数
	// counter[num] 代表 num 的出现次数
	counter := make([]int, m+1)
	for _, num := range nums {
		counter[num]++
	}
	// 3. 求 counter 的前缀和，将“出现次数”转换为“尾索引”
	// 即 counter[num]-1 是 num 在 res 中最后一次出现的索引
	for i := 0; i < m; i++ {
		counter[i+1] += counter[i]
	}
	// 4. 倒序遍历 nums ，将各元素填入结果数组 res
	// 初始化数组 res 用于记录结果
	n := len(nums)
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		// 将 num 放置到对应索引处
		res[counter[num]-1] = num
		// 令前缀和自减 1 ，得到下次放置 num 的索引
		counter[num]--
	}
	// 使用结果数组 res 覆盖原数组 nums
	copy(nums, res)
}

/* 获取元素 num 的第 k 位，其中 exp = 10^(k-1) */
func digit(num, exp int) int {
	// 传入 exp 而非 k 可以避免在此重复执行昂贵的次方计算
	return (num / exp) % 10
}

/* 计数排序（根据 nums 第 k 位排序） */
func countingSortDigit(nums []int, exp int) {
	// 十进制的位范围为 0~9 ，因此需要长度为 10 的桶数组
	counter := make([]int, 10)
	n := len(nums)
	// 统计 0~9 各数字的出现次数
	for i := 0; i < n; i++ {
		d := digit(nums[i], exp) // 获取 nums[i] 第 k 位，记为 d
		counter[d]++             // 统计数字 d 的出现次数
	}
	// 求前缀和，将“出现个数”转换为“数组索引”
	for i := 1; i < 10; i++ {
		counter[i] += counter[i-1]
	}
	// 倒序遍历，根据桶内统计结果，将各元素填入 res
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		d := digit(nums[i], exp)
		j := counter[d] - 1 // 获取 d 在数组中的索引 j
		res[j] = nums[i]    // 将当前元素填入索引 j
		counter[d]--        // 将 d 的数量减 1
	}
	// 使用结果覆盖原数组 nums
	for i := 0; i < n; i++ {
		nums[i] = res[i]
	}
}

/* 基数排序 */
func radixSort(nums []int) {
	// 获取数组的最大元素，用于判断最大位数
	max := math.MinInt
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	// 按照从低位到高位的顺序遍历
	for exp := 1; max >= exp; exp *= 10 {
		// 对数组元素的第 k 位执行计数排序
		// k = 1 -> exp = 1
		// k = 2 -> exp = 10
		// 即 exp = 10^(k-1)
		countingSortDigit(nums, exp)
	}
}
