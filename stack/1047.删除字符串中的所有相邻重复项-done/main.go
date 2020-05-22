/**
 * @Time : 2020-05-15 10:36
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-leetcode/types"
)

func removeDuplicates(S string) string {
	s := types.ConstructStack()
	if len(S) > 1 {
		s.Push(fmt.Sprintf("%c", S[0]))
	} else {
		return S
	}
	for i := 1; i < len(S); i++ {
		temp := fmt.Sprintf("%c", S[i])
		if s.Top() != nil && temp == s.Top().(string) {
			s.Pop()
		} else {
			s.Push(temp)
		}
	}
	var r string
	for i := 0 ; i < len(s.Elem); i++ {
		r += s.Elem[i].(string)
	}
	return r
}

func main()  {
	fmt.Println(removeDuplicates("abbcc"))
}