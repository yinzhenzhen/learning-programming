/**
 * @Time : 2020-09-11 19:14
 * @Author : yz
 */

package string

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
标签：字符串
难度：简单

题目内容：
HJ1.字符串最后一个单词的长度
计算字符串最后一个单词的长度，单词以空格隔开，字符串长度小于5000。（注：字符串末尾不以空格为结尾）

解题思路：有没有更好的解法？

*/

func HJ1() {
	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		if input == "" {
			return
		}
		input = input[:len(input)-1]
		fmt.Println(lastWordLength(input))
	}
}

func lastWordLength(str string) int {
	if str == "" {
		return 0
	}
	words := strings.Split(str, " ")
	return len(words[len(words)-1])
}
