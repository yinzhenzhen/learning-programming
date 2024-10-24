/**
 * @Time : 2020-05-31 22:23
 * @Author : yz
 */

/**
题目内容：
56.合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

解题思路：

*/

package sort

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var newIntervals = make([][]int, len(intervals))
	for i := 0; i < len(intervals)-1; i++ {
		temp := make([]int, 2)
		if intervals[i][1] > intervals[i+1][0] {
			temp[0] = intervals[i][0]
			if intervals[i][1] > intervals[i+1][1] {
				temp[1] = intervals[i][1]
			} else {
				temp[1] = intervals[i+1][1]
			}
		} else {
			temp[0] = intervals[i][0]
			temp[1] = intervals[i][1]
		}
		newIntervals[i] = append(newIntervals[i], temp...)
	}

	return newIntervals
}
