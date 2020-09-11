/**
 * @Time : 2020-09-11 20:50
 * @Author : yz
 */

package main

import (
	"fmt"
)

func split(input string) (output []string) {
	if len(input) == 8 {

		output = append(output, input)
		return

	}

	for len(input) > 8 {

		temp := input[0:8]
		output = append(output, temp)

		input = input[8:]
	}

	add0 := 8 - len(input)

	for i := 0; i < add0; i++ {
		input = input + "0"
	}

	output = append(output, input)
	return
}

func main()  {
	var input1, input2 string

	_, _ = fmt.Scan(&input1, &input2)

	output1 := split(input1)
	output2 := split(input2)

	for _, v := range output1 {
		fmt.Println(v)
	}

	for _, v := range output2 {
		fmt.Println(v)
	}

}