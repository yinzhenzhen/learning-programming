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

题目描述：
HJ106.字符逆序
将一个字符串str的内容颠倒过来，并输出。数据范围：1≤len(str)≤10000

解题思路：

*/

func H106() {
	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		if input == "" {
			return
		}
		input = input[:len(input)-1]
		fmt.Println(characterReverseOrder(input))
	}
}

func characterReverseOrder(s string) string {
	sslice := strings.Split(s, "")
	i, j := 0, len(sslice)-1
	for i <= j {
		swap(&sslice[i], &sslice[j])
		i++
		j--
	}
	return strings.Join(sslice, "")
}

func swap(a, b *string) {
	temp := *a
	*a = *b
	*b = temp
}
