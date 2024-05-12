/**
 * @Time : 4/12/22 4:23 PM
 * @Author : yz
 */

package main

import "fmt"

/**
单词长度的最大乘积
*/

func main() {
	fmt.Println(maxProduct([]string{"abcw", "baz", "foo", "bar", "fxyz", "abcdef"}))
	fmt.Println(maxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}))
	fmt.Println(maxProduct([]string{"a", "aa", "aaa", "aaaa"}))
	fmt.Println(maxProduct([]string{"eae", "ea", "aaf", "bda", "fcf", "dc", "ac", "ce", "cefde", "dabae"}))
}

// 优化：https://leetcode-cn.com/problems/aseY1I/solution/jian-dan-yi-dong-javac-pythonjs-zui-da-d-ffga/

type Word struct {
	letter map[string]int
	len    int
}

func maxProduct(words []string) int {

	w := make([]Word, 0)

	for i := 0; i < len(words); i++ {
		y := strToMap(words[i])
		w = append(w, Word{
			y,
			len(words[i]),
		})
	}

	max := 0
	for _, i := range w {
		for _, j := range w {

			if !isRepeat(i, j) {
				temp := i.len * j.len
				if temp >= max {
					max = temp
				}
			}

		}
	}

	return max
}

func isRepeat(a Word, b Word) bool {

	for k, _ := range a.letter {
		if _, ok := b.letter[k]; ok {
			return true
		}
	}
	return false
}

func strToMap(word string) (r map[string]int) {
	r = make(map[string]int)
	for _, v := range word {
		if _, ok := r[string(rune(v))]; ok {
			r[string(rune(v))]++
		} else {
			r[string(rune(v))] = 1
		}

	}
	return r
}
