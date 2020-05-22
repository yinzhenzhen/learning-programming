/**
 * @Time : 2020-05-22 10:09
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-leetcode/types"
	"strconv"
)

func evalRPN(tokens []string) int {
	var result int
	var newStack = types.ConstructStack()

	for _, v := range tokens {
		if v == "+" {
			num1, _ := strconv.Atoi(newStack.Pop().(string))
			num2, _ := strconv.Atoi(newStack.Pop().(string))
			newStack.Push(strconv.Itoa(num2 + num1))

		} else if v == "-" {
			num1, _ := strconv.Atoi(newStack.Pop().(string))
			num2, _ := strconv.Atoi(newStack.Pop().(string))
			newStack.Push(strconv.Itoa(num2 - num1))

		} else if v == "*" {
			num1, _ := strconv.Atoi(newStack.Pop().(string))
			num2, _ := strconv.Atoi(newStack.Pop().(string))
			newStack.Push(strconv.Itoa(num2 * num1))

		} else if v == "/" {
			num1, _ := strconv.Atoi(newStack.Pop().(string))
			num2, _ := strconv.Atoi(newStack.Pop().(string))
			newStack.Push(strconv.Itoa(num2 / num1))

		} else {
			newStack.Push(v)
		}
	}

	result, _ = strconv.Atoi(newStack.Top().(string))

	return result
}

func main()  {
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}