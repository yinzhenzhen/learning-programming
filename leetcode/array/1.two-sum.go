package array

/**
题目内容：
1.两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值 target 的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。

解题思路：
引入map结构，查找两数之和中的另一个元素，另一个元素=target-当前遍历到的元素，
如果另一个元素不在map中，则把当前遍历的元素放入map（如果有与之匹配的元素，之后的元素查找另一半可以在map中直接获取到），
如果另一个元素在map中，则说明满足条件的是当前遍历到的元素和map中已经缓存的元素
*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		other := target - v
		if j, ok := m[other]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}
