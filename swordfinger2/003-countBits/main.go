/**
 * @Time : 4/12/22 1:42 PM
 * @Author : yz
 */

package main

import "fmt"

func main() {
	fmt.Println(countBits(10))
}

/*
*
前n个数字二进制中1的个数
*/
func countBits(n int) []int {

	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}
	if n == 2 {
		return []int{0, 1, 1}
	}

	r := make([]int, n+1)
	r[0] = 0
	r[1] = 1
	r[2] = 1
	j := 2
	for i := 3; i < n+1; {

		if i == j || i == j*2 {
			r[i] = 1
			//fmt.Println(r[i])
			i++
		} else if i > j && i < j*2 {
			r[i] = r[i-j] + 1
			//fmt.Println(r[i])
			i++
		} else if i > j*2 {
			j = j * 2
		}
	}

	return r
}
