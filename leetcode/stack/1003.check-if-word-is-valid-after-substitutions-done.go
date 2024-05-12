/**
 * @Time : 2020-05-23 00:00
 * @Author : yz
 */

package stack

import (
	"fmt"
	"github.com/yinzhenzhen/learning-programming/types"
)

func isValid(S string) bool {
	if len(S) < 3 {
		return false
	}

	newStack := types.ConstructStack()

	a := fmt.Sprintf("%c", S[0])
	if a != "a" {
		return false
	}

	newStack.Push(a)

	for i := 1; i < len(S); i++ {
		str := fmt.Sprintf("%c", S[i])

		if str == "c" {
			if newStack.Size() < 2 {
				return false
			}
			num1 := newStack.Pop().(string)
			if num1 != "b" {
				newStack.Push(num1)
			} else {
				num2 := newStack.Pop().(string)
				if num2 != "a" {
					newStack.Push(num2)
				}
			}
		} else {
			newStack.Push(str)
		}
	}

	if newStack.Size() == 0 {
		return true
	}
	return false
}
