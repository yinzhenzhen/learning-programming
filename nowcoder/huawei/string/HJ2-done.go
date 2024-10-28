/**
 * @Time : 2020-09-11 19:36
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
标签：字符串、哈希
难度：简单

题目内容：
HJ2.计算某字符出现次数
写出一个程序，接受一个由字母、数字和空格组成的字符串，和一个字符，然后输出输入字符串中该字符的出现次数。（不区分大小写字母）
数据范围：1≤n≤1000

解题思路：

*/

func HJ2() {
	for {
		var str, single string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		str = scanner.Text()
		if str == "" {
			return
		}
		scanner.Scan()
		single = scanner.Text()
		fmt.Println(wordCount2(str, single))
	}
}

func wordCount(str string, single string) (count int) {
	for _, v := range str {
		if string(v) == single || string(v-32) == single || string(v+32) == single {
			count++
		}
	}
	return
}

// 借助 map
func wordCount2(str string, single string) (count int) {
	m := make(map[string]int)
	for _, k := range str {
		if string(k) == " " {
			continue
		}
		kk := strings.ToUpper(string(k))
		if v, ok := m[kk]; ok {
			m[kk] = v + 1
		} else {
			m[kk] = 1
		}
	}
	singleUpper := strings.ToUpper(single)
	return m[singleUpper]
}
