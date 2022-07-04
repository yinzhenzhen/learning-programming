/**
 * @Time : 2022/7/3 3:18 下午
 * @Author : yz
 */

package types

// BinaryIndexedTree define
// 树状数组中父子节点下标关系是 parent = son + 2^k, k 是子节点下标对应二进制末尾 0 的个数
type BinaryIndexedTree struct {
	tree     []int
	capacity int
}

func (bit *BinaryIndexedTree) Init(capacity int) {
	bit.tree, bit.capacity = make([]int, capacity+1), capacity
}

// Init define
func (bit *BinaryIndexedTree) InitWithNums(nums []int) {
	bit.tree, bit.capacity = make([]int, len(nums)+1), len(nums)+1
	for i := 1; i <= len(nums); i++ {
		bit.tree[i] += nums[i-1]
		for j := i - 2; j >= i-lowbit(i); j-- {
			bit.tree[i] += nums[j]
		}
	}
}

// lowbit(i) 函数返回 i 转换成二进制以后，末尾最后一个 1 代表的数值，即 2^k，k 为 i 末尾 0 的个数
func lowbit(x int) int {
	return x & -x
}

// Add define
func (bit *BinaryIndexedTree) Add(index int, val int) {
	for index < bit.capacity {
		bit.tree[index] += val
		index += lowbit(index)
	}
}

// Query define
func (bit *BinaryIndexedTree) Query(index int) int {
	sum := 0
	for index >= 1 {
		sum += bit.tree[index]
		index -= lowbit(index)
	}
	return sum
}