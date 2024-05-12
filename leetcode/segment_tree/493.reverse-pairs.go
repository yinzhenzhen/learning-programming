package segment_tree

import (
	"sort"

	"github.com/yinzhenzhen/learning-programming/types"
)

func reversePairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	st, numsMap, indexMap, numsArray, res := types.SegmentCountTree{}, make(map[int]int, 0), make(map[int]int, 0), []int{}, 0
	numsMap[nums[0]] = nums[0]
	for _, num := range nums {
		numsMap[num] = num
		numsMap[2*num+1] = 2*num + 1
	}
	// numsArray 是 prefixSum 去重之后的版本，利用 numsMap 去重
	for _, v := range numsMap {
		numsArray = append(numsArray, v)
	}
	// 排序是为了使得线段树中的区间 left <= right，如果此处不排序，线段树中的区间有很多不合法。
	sort.Ints(numsArray)
	// 离散化，构建映射关系
	for i, n := range numsArray {
		indexMap[n] = i
	}
	numsArray = []int{}
	// 离散化，此题如果不离散化，MaxInt32 的数据会使得数字越界。
	for i := 0; i < len(indexMap); i++ {
		numsArray = append(numsArray, i)
	}
	// 初始化线段树，节点内的值都赋值为 0，即计数为 0
	st.Init(numsArray, func(i, j int) int {
		return 0
	})
	for _, num := range nums {
		res += st.Query(indexMap[num*2+1], len(indexMap)-1)
		st.UpdateCount(indexMap[num])
	}
	return res
}
