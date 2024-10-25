/**
 * @Time : 2020-05-22 10:09
 * @Author : yz
 */

package stack

import (
	"github.com/yinzhenzhen/learning-programming/types"
	"strconv"
)

/**
标签：栈、数组、数学
难度：中等

题目内容：
150.逆波兰表达式求值
给你一个字符串数组 tokens，表示一个根据逆波兰表示法表示的算术表达式。
请你计算该表达式。返回一个表示表达式值的整数。
注意：
	有效的算符为 '+'、'-'、'*' 和 '/' 。
	每个操作数（运算对象）都可以是一个整数或者另一个表达式。
	两个整数之间的除法总是 向零截断 。
	表达式中不含除零运算。
	输入是一个根据逆波兰表示法表示的算术表达式。
	答案及所有中间计算结果可以用 32 位整数表示。

解题思路：见代码注释

逆波兰表达式：
逆波兰表达式是一种后缀表达式，所谓后缀就是指算符写在后面。
	平常使用的算式则是一种中缀表达式，如 ( 1 + 2 ) * ( 3 + 4 ) 。
	该算式的逆波兰表达式写法为 ( ( 1 2 + ) ( 3 4 + ) * ) 。
逆波兰表达式主要有以下两个优点：
	去掉括号后表达式无歧义，上式即便写成 1 2 + 3 4 + * 也可以依据次序计算出正确结果。
	适合用栈操作运算：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中
*/

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
