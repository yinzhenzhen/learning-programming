/**
 * @Time : 2020-09-15 20:03
 * @Author : yz
 */

package huawei

import "fmt"

func game(M int, nums []int) (result []int) {

	// 新一轮的开始报数的位置
	newIndex := 0

	for len(nums) >= M {

		// 本轮会退出的个数
		exitNum := len(nums) / M

		nums, newIndex = handle(nums, newIndex, exitNum, M)
	}

	result = nums

	return
}

// 每次处理相应的数据
func handle(nums []int, newIndex int, exitNum int, M int) (result []int, nextNewIndex int) {

	//fmt.Println(nums)
	//fmt.Println(len(nums))
	//fmt.Println(newIndex)

	exit := []int{}

	start := 1

	nextNum := 0

	// 确定要被删除的元素
	for j := newIndex; len(exit) != exitNum; j++ {

		// 如果当前索引超出长度，跳转到第一个元素
		if j >= len(nums) {
			j = 0
		}
		// 如果该元素要被删除
		if start == M {
			exit = append(exit, nums[j])

			start = 0
		}

		// 如果本轮要删除的元素全部删除
		if len(exit) == exitNum {
			// 记录下一轮的id
			if j == len(nums)-1 {
				nextNum = nums[0]
			} else {
				nextNum = nums[j+1]
			}

			break
		}

		start++
	}

	//fmt.Println(exit)
	//fmt.Println(nextNum)
	//fmt.Println("-----------------------------")

	for j := 0; j < len(exit); j++ {

		for i := 0; i < len(nums); i++ {

			if nums[i] == exit[j] {

				temp := nums[0:i]
				temp = append(temp, nums[i+1:]...)
				nums = temp
				break
			}
		}

	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == nextNum {
			nextNewIndex = i
			break
		}
	}

	result = nums

	return
}

func T2() {
	//var M int
	//
	//for {
	//	n, _ := fmt.Scan(&M)
	//	if n == 0 {
	//		break
	//	}
	//
	//
	//	nums := []int{}
	//
	//	for i := 0; i <= 100; i++ {
	//		nums = append(nums, i+1)
	//	}
	//
	//	fmt.Println(game(M, nums))
	//}

	nums := []int{}

	for i := 0; i < 100; i++ {
		nums = append(nums, i+1)
	}

	result := game(4, nums)

	for i, v := range result {
		if i == len(result)-1 {
			fmt.Print(v)
		} else {
			fmt.Print(v, ",")
		}
	}
}
