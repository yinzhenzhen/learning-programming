/**
 * @Time : 2020-05-15 10:55
 * @Author : yz
 */

package stack

/**
有效的括号：
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
1、左括号必须用相同类型的右括号闭合。
2、左括号必须以正确的顺序闭合。
3、每个右括号都有一个对应的相同类型的左括号。

解题思路：

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
