/**
 * @Time : 2020-06-05 17:12
 * @Author : yz
 */

package main

import (
	"fmt"
)

func nthPersonGetsNthSeat(n int) float64 {
	return 0.0
}

func main()  {
	//fmt.Println(nthPersonGetsNthSeat(3))
	s := []int{1, 3, 5, 7}
	for i, v := range s  {
		if v == 7 {
			s = append(s[:i], s[i+1:]...)
		}
	}
	fmt.Println(s)
}