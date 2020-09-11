/**
 * @Time : 2020-05-15 15:08
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-programming/types"
)

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
			start = i+1
			continue
		// 如果新元素是/，标记当前位置为目录结束位置
		} else if (str == "/" && end == 0) {
			end = i
		// 最后一个元素，但是不是以/结尾
		} else if str != "/" && i == len(path) - 1 {
			end = len(path)
		} else {
			continue;
		}

		if start != 0 && end != 0 {
			dir = path[start : end]
			start = 0
			end = 0
			if i != len(path) - 1 {
				i--
			}
		}

		// 如果是当前目录,或者//
		if dir == "." || dir == "" {
			continue;
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

func main()  {
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
}