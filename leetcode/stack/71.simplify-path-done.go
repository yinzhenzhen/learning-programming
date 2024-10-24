/**
 * @Time : 2020-05-15 15:08
 * @Author : yz
 */

package stack

import (
	"fmt"
	"github.com/yinzhenzhen/learning-programming/types"
)

/**
标签：栈、字符串
难度：简单

题目内容：
71.简化路径：
给你一个字符串 path ，表示指向某一文件或目录的 Unix 风格 绝对路径 （以 '/' 开头），请你将其转化为更加简洁的规范路径。
在 Unix 风格的文件系统中规则如下：
一个点 '.' 表示当前目录本身。
此外，两个点 '..' 表示将目录切换到上一级（指向父目录）。
任意多个连续的斜杠（即，'//' 或 '///'）都被视为单个斜杠 '/'。
任何其他格式的点（例如，'...' 或 '....'）均被视为有效的文件/目录名称。
返回的简化路径必须遵循下述格式：
始终以斜杠 '/' 开头。
两个目录名之间必须只有一个斜杠 '/' 。
最后一个目录名（如果存在）不能 以 '/' 结尾。
此外，路径仅包含从根目录到目标文件或目录的路径上的目录（即，不含 '.' 或 '..'）。
返回简化后得到的规范路径。

解题思路：见代码注释

*/

func simplifyPath(path string) string {
	if path == "" || path == "/" {
		return path
	}

	newStack := types.ConstructStack()
	var dir string
	var start, end int = 0, 0

	for i := 0; i < len(path); i++ {
		str := fmt.Sprintf("%c", path[i])

		// 如果新元素是/，标记当前位置为目录开始位置，寻找下一个/
		if str == "/" && start == 0 {
			start = i + 1
			continue
			// 如果新元素是/，标记当前位置为目录结束位置
		} else if str == "/" && end == 0 {
			end = i
			// 最后一个元素，但是不是以/结尾
		} else if str != "/" && i == len(path)-1 {
			end = len(path)
		} else {
			continue
		}

		if start != 0 && end != 0 {
			dir = path[start:end]
			start = 0
			end = 0
			if i != len(path)-1 {
				i--
			}
		}

		// 如果是当前目录,或者//
		if dir == "." || dir == "" {
			continue
			// 如果是上一级目录
		} else if dir == ".." {
			newStack.Pop()
			// 如果是目录名称
		} else {
			newStack.Push(dir)
		}

	}

	var result string
	for _, v := range newStack.Elem[:] {
		if v.(string) != "" && v.(string) != "/" {
			result += "/" + v.(string)
		}
	}

	if result == "" {
		return "/"
	}

	return result
}
