/**
 * @Time : 2020-09-11 19:59
 * @Author : yz
 */

package array

import (
	"fmt"
	"sort"
)

/**
标签：数组
难度：较难

题目内容：
HJ3.明明的随机数
明明生成了N个1到500之间的随机整数。请你删去其中重复的数字，即相同的数字只保留一个，把其余相同的数去掉，然后再把这些数从小到大排序，按照排好的顺序输出。
数据范围：1≤n≤1000，输入的数字大小满足 1≤val≤500

解题思路：
先排序，然后遍历字符数组，如果当前字符和前一个一致，则删除该字符
*/

func random() {
	var num int

	for {
		var input []int
		n, _ := fmt.Scan(&num)
		if n == 0 {
			return
		}
		for i := 0; i < num; i++ {
			var temp int
			_, _ = fmt.Scan(&temp)
			input = append(input, temp)
		}
		output := randomInternal(input)
		for _, v := range output {
			fmt.Println(v)
		}
	}
}

func randomInternal(input []int) (output []int) {
	if len(input) == 0 {
		return
	}

	sort.Ints(input)

	output = []int{input[0]}

	for i := 1; i < len(input); i++ {
		if input[i] == output[i-1] {
			var temp []int
			temp = append(temp, input[0:i]...)
			temp = append(temp, input[i+1:]...)
			input = temp
			i--
		} else {
			output = append(output, input[i])
		}
	}
	return
}
