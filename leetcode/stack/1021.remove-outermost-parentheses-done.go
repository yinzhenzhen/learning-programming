/**
 * @Time : 2020-05-14 17:33
 * @Author : yz
 */

package stack

import (
	"fmt"
	"github.com/yinzhenzhen/learning-programming/types"
	"strings"
)

func removeOuterParentheses(S string) string {
	var primitive []string
	in := types.ConstructStack()
	out := types.ConstructStack()
	num1 := 0
	num2 := 0
	for i := 0; i < len(S); i++ {
		temp := fmt.Sprintf("%c", S[i])
		in.Push(temp)
		if temp == "(" {
			num1++
		} else if temp == ")" {
			num2++
		}
		if num1 == num2 {
			// 删除栈顶元素
			in.Pop()
			for in.Size() > 1 {
				out.Push(in.Pop())
			}
			// 删除栈底元素
			in.Pop()
			var s string
			for out.Size() > 0 {
				s += out.Pop().(string)
			}
			primitive = append(primitive, s)
			num1, num2 = 0, 0
		}
	}

	return strings.Join(primitive, "")
}
