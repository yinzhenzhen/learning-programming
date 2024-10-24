/**
 * @Time : 2020-05-15 10:55
 * @Author : yz
 */

package stack

/**
标签：栈、字符串
难度：简单

题目内容：
20.有效的括号：
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s，判断字符串是否有效。
有效字符串需满足：
1、左括号必须用相同类型的右括号闭合。
2、左括号必须以正确的顺序闭合。
3、每个右括号都有一个对应的相同类型的左括号。

解题思路：见代码注释

*/

import (
	"fmt"
	"github.com/yinzhenzhen/learning-programming/types"
)

func isValidParentheses(s string) bool {
	newStack := types.ConstructStack()
	if s == "" {
		return true
	}
	if len(s) <= 1 {
		return false
	}
	// 向栈中写入第一个元素
	newStack.Push(fmt.Sprintf("%c", s[0]))
	for i := 1; i < len(s); i++ {
		temp := fmt.Sprintf("%c", s[i])
		if newStack.Top() == nil {
			// 如果栈为空，则写入元素
			newStack.Push(temp)
			continue
		}
		// 获得栈顶元素
		top := newStack.Top().(string)
		// 如果当前元素和栈顶元素匹配，则从栈中取出栈顶元素
		if (top == "(" && temp == ")") || (top == "[" && temp == "]") || (top == "{" && temp == "}") {
			newStack.Pop()
		} else {
			newStack.Push(temp)
		}
	}

	// 如果最终的栈为空，则说明元素括号都是匹配的
	if !newStack.IsEmpty() {
		return false
	}

	return true
}
