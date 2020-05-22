/**
 * @Time : 2020-05-14 17:33
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-leetcode/types"
	"strings"
)



func removeOuterParentheses(S string) string {
	var primitive []string
	in := types.Construct()
	out := types.Construct()
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



func main()  {
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
	//s := stack.Construct()
	//s.Push(1)
	//s.Push(2)
	//s.Push(3)
	//fmt.Println(s.Pop())
	//fmt.Println(s.Pop())
	//fmt.Println(s.Pop())
}

