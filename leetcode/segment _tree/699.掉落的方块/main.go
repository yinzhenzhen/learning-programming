package main

import (
	"fmt"
	"sort"

	"github.com/yinzhenzhen/learning-leetcode/types"
)

func fallingSquares(positions [][]int) []int {
	st, ans, posMap, maxHeight := types.SegmentTree{}, make([]int, 0, len(positions)), discretization(positions), 0
	tmp := make([]int, len(posMap))
	st.Init(tmp, func(i, j int) int {
		return max(i, j)
	})
	for _, p := range positions {
		h := st.QueryLazy(posMap[p[0]], posMap[p[0]+p[1]-1]) + p[1]
		st.UpdateLazy(posMap[p[0]], posMap[p[0]+p[1]-1], h)
		maxHeight = max(maxHeight, h)
		ans = append(ans, maxHeight)
	}
	return ans
}

func discretization(positions [][]int) map[int]int {
	tmpMap, posArray, posMap := map[int]int{}, []int{}, map[int]int{}
	for _, pos := range positions {
		tmpMap[pos[0]]++
		tmpMap[pos[0]+pos[1]-1]++
	}
	for k := range tmpMap {
		posArray = append(posArray, k)
	}
	sort.Ints(posArray)
	for i, pos := range posArray {
		posMap[pos] = i
	}
	return posMap
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(fallingSquares([][]int{{1,2}, {2,3}, {6,1}}))
}
