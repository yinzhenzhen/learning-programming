/**
 * @Time : 4/11/22 5:01 PM
 * @Author : yz
 */

package main

import (
	"fmt"
	"math"
)

// 超时
func main() {

	fmt.Println(divide(-2147483648, 2))
}

func divide(a int, b int) int {

	if a == math.MinInt32 { // 考虑被除数为最小值的情况
		if b == 1 {
			return math.MinInt32
		}
		if b == -1 {
			return math.MaxInt32
		}
	}
	if b == math.MinInt32 { // 考虑除数为最小值的情况
		if a == math.MinInt32 {
			return 1
		}
		return 0
	}

	if a == 0 {
		return 0
	}

	if a == b {
		return 1
	} else if a == -b {
		return -1
	}

	if b == 1 {
		return a
	} else if b == -1 {
		return -a
	}

	isNeg := false

	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		if a < 0 {
			a = -a
		} else if b < 0 {
			b = -b
		}
		isNeg = true
	} else if a < 0 && b < 0 {
		a = -a
		b = -b
	}

	if a < b {
		return 0
	}

	result := 0
	for a >= b {
		a = a - b
		result++
	}

	if isNeg {
		result = -result
	}

	return result
}
