/**
 * @Time : 2020-05-25 16:06
 * @Author : yz
 */

package stack

import (
	"fmt"
	"github.com/yinzhenzhen/learning-programming/types"
)

func validateStackSequences(pushed []int, popped []int) bool {
	newStack := types.ConstructStack()

	pushIndex := 0
	popIndex := 0

	for i := 0; i < len(pushed); i++ {
		for newStack.Top() != nil && popped[popIndex] == newStack.Top().(int) {
			x := newStack.Pop()
			fmt.Println("pop ", x)
			popIndex++
		}

		newStack.Push(pushed[pushIndex])
		fmt.Println("push ", pushed[pushIndex])
		pushIndex++

	}

	for !newStack.IsEmpty() && popIndex < len(popped) {
		if newStack.Top() != nil && newStack.Top() == popped[popIndex] {
			x := newStack.Pop()
			fmt.Println("pop ", x)
			popIndex++
		} else {
			return false
		}
	}

	if newStack.Size() == 0 {
		return true
	}

	return false
}
