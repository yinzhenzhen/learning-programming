/**
 * @Time : 2020-09-11 19:59
 * @Author : yz
 */

package main

import (
	"fmt"
	"sort"
)

func random(input []int) (output []int) {
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

func main()  {
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

		output := random(input)
		for _, v := range output{
			fmt.Println(v)
		}

	}
}