package recursion

import (
	"fmt"
)

/**
题目内容：
最小公倍数：正整数A和正整数B的最小公倍数是指能被A和B整除的最小的正整数值，设计一个算法，求输入A和B的最小公倍数。

解题思路：

*/

func leastCommonMultiple() {
	a := 0
	b := 0
	for {
		n, _ := fmt.Scan(&a, &b)
		if n == 0 {
			break
		} else {
			c := lcm(a, b)
			fmt.Printf("%d\n", c)
		}
	}
}

func lcm(a, b int) int {
	// 最小公倍数为两数之积除以最大公约数
	return a * b / gcd(a, b)
}

// 辗转相除法求最大公约数
func gcd(a, b int) int {
	if a < b {
		temp := a
		a = b
		b = temp
	}
	for b > 0 {
		remain := a % b
		a = b
		b = remain
	}
	return a
}
