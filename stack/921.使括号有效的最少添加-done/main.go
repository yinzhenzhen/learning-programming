/**
 * @Time : 2020-05-25 16:39
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-leetcode/types"
)

func minAddToMakeValid(S string) int {
	if S == "" {
		return 0
	}
	newStack := types.ConstructStack()

	for i := 0; i < len(S); i++ {
		str := fmt.Sprintf("%c", S[i])
		if str == "(" {
			newStack.Push(str)
			fmt.Println("push ", str)
		} else {
			if newStack.Top() == "(" {
				e := newStack.Pop()
				fmt.Println("pop ", e)
			} else {
				newStack.Push(str)
				fmt.Println("push ", str)
			}
		}
	}

	return newStack.Size()
}

func main()  {
	fmt.Println(minAddToMakeValid("()"))
}