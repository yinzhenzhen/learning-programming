/**
 * @Time : 2020-05-15 10:55
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-programming/types"
)

func isValid(s string) bool {
	newStack := types.ConstructStack()
	if s == "" {
		return true
	}
	if len(s) <= 1 {
		return false
	}
	newStack.Push(fmt.Sprintf("%c", s[0]))
	for i := 1; i < len(s); i++ {
		temp := fmt.Sprintf("%c", s[i])
		if newStack.Top() == nil {
			newStack.Push(temp)
			continue
		}
		top := newStack.Top().(string)
		if (top == "(" && temp == ")") || (top == "[" && temp == "]") || (top == "{" && temp == "}") {
			newStack.Pop()
		} else {
			newStack.Push(temp)
		}
	}

	if !newStack.IsEmpty() {
		return false
	}

	return true
}

func main()  {
	fmt.Println(isValid("[]{"))
}