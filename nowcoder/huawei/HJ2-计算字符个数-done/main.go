/**
 * @Time : 2020-09-11 19:36
 * @Author : yz
 */

package main

import (
	"fmt"
)

func wordCount(str string, single string) (count int) {
	for _, v := range str {
		if string(v) == single || string(v-32) == single || string(v+32) == single {
			count++
		}
	}
	return
}

func main()  {
	var str string
	var single string
	
	for {
		n, _ := fmt.Scan(&str, &single)
		if n == 0 {
			return
		}
		
		fmt.Println(wordCount(str, single))
	}
}