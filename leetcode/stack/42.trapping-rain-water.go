/**
 * @Time : 2020-05-15 13:01
 * @Author : yz
 */

package stack

/**
接雨水：

解题思路：

*/

import (
	"github.com/yinzhenzhen/learning-programming/types"
)

func trap(height []int) int {
	var result int
	stack1 := types.ConstructStack()
	if len(height) <= 1 {
		return 0
	}
	stack1.Push(height[0])
	for i := 1; i < len(height); i++ {
		if stack1.Top().(int) > height[i] && i != len(height)-1 {
			stack1.Push(stack1.Top())
		} else {
			stack1.Push(height[i])
		}

		result += stack1.Top().(int) - height[i]
	}

	return result
}
