/**
 * @Time : 2020-06-02 20:18
 * @Author : yz
 */

package sort

import (
	"fmt"
	"strings"
)

type Word struct {
	content string
	length  int
}

// 冒泡排序
func arrangeWords(text string) string {
	if len(text) == 1 {
		return text
	}

	// 单词起始位置
	start := 0
	// 所有单词
	words := []Word{}
	// 单个单词的所有字母
	singleWord := []string{}

	for i := 0; i < len(text); i++ {
		temp := fmt.Sprintf("%c", text[i])
		// 如果到了空格，单词结束
		if temp == " " {
			// 构建单个单词
			word := Word{strings.ToLower(strings.Join(singleWord, "")), i - start}
			// 添加到单词切片
			words = append(words, word)
			// 新的单词为空
			singleWord = []string{}
			start = i + 1
			// 单词的字母切片
		} else {
			singleWord = append(singleWord, temp)
		}

		if i == len(text)-1 {
			word := Word{strings.ToLower(strings.Join(singleWord, "")), i - start + 1}
			words = append(words, word)
		}
	}

	// 冒泡排序
	flag := true
	for i := 0; i < len(words) && flag; i++ {
		flag = false
		for j := len(words) - 2; j >= i; j-- {
			word1 := words[j]
			word2 := words[j+1]
			if word1.length > word2.length {
				swap(words, j, j+1)
				flag = true
			}
		}
	}

	for i := 0; i < len(words); i++ {
		fmt.Println(words[i])

	}

	result := []string{}
	// 将第一个单词的第一个字母转为大写
	first := strings.Split(words[0].content, "")
	first[0] = capitalize(first[0])
	f := strings.Join(first, "")
	result = append(result, f)

	for i := 1; i < len(words); i++ {
		result = append(result, " ", words[i].content)
	}

	return strings.Join(result, "")
}

// 交换
func swap(word []Word, i int, j int) {
	temp := word[i]
	word[i] = word[j]
	word[j] = temp
}

// 首字母大写
func capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
