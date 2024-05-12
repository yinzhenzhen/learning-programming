/**
 * @Time : 2020-05-28 15:21
 * @Author : yz
 */

package stack

import (
	"fmt"
	"github.com/yinzhenzhen/learning-programming/types"
)

func backspaceCompare(S string, T string) bool {
	s := types.ConstructStack()
	t := types.ConstructStack()

	length := 0

	if len(S) > len(T) {
		length = len(S)
	} else {
		length = len(T)
	}

	for i := 0; i < length; i++ {
		if i < len(S) {
			str := fmt.Sprintf("%c", S[i])
			if str == "#" && !s.IsEmpty() {
				s.Pop()
			} else if str != "#" {
				s.Push(S[i])
			}
		}

		if i < len(T) {
			str := fmt.Sprintf("%c", T[i])
			if str == "#" && !t.IsEmpty() {
				t.Pop()
			} else if str != "#" {
				t.Push(T[i])
			}
		}
	}

	if s.Size() != t.Size() {
		return false
	}

	for i := 0; !s.IsEmpty(); i++ {
		ss := s.Pop()
		tt := t.Pop()
		if ss != tt {
			return false
		}
	}

	return true

}
