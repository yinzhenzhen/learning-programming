/**
 * @Time : 2020-09-14 22:44
 * @Author : yz
 */

package main

import (
	"fmt"
	"strconv"
)

var f = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}


func baseConversion(str string) (result int) {

	var all string

	for _, v := range str  {
		all += f[string(v)]
	}

	pos := 1

	for i := len(all)-1; i >= 0; i-- {

		temp, _ := strconv.Atoi(string(all[i]))

		result += pos * temp

		pos = pos * 2
	}

	return
}

func main()  {
	var data string
	for {
		n, _ := fmt.Scan(&data)
		if n == 0 {
			break
		}
		
		fmt.Println(baseConversion(data[2:]))
	}
}
