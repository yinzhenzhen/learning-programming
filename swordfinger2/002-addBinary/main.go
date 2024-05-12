/**
 * @Time : 4/11/22 5:07 PM
 * @Author : yz
 */

package main

import (
	"fmt"
)

func main() {

	fmt.Println(addBinary("11", "10"))
	fmt.Println(addBinary("1111", "1111"))
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("1", "111"))
}

func addBinary(a string, b string) string {

	if a == "0" {
		return b
	}
	if b == "0" {
		return a
	}

	sta := ConstructStack()
	stb := ConstructStack()
	for _, i := range a {
		sta.Push(string(rune(i)))
	}

	for _, i := range b {
		stb.Push(string(rune(i)))
	}

	result := ConstructStack()
	flag := false
	for {
		temp1 := sta.Pop()
		temp2 := stb.Pop()
		if temp1 == "1" && temp2 == "1" {
			if flag {
				result.Push("1")
			} else {
				result.Push("0")
			}
			flag = true

		} else if temp1 == "0" && temp2 == "0" {
			if flag {
				result.Push("1")
			} else {
				result.Push("0")
			}
			flag = false

		} else {
			if flag {
				result.Push("0")
				flag = true
			} else {
				result.Push("1")
				flag = false
			}
		}
		if sta.IsEmpty() && stb.IsEmpty() {
			if flag {
				result.Push("1")
			}
			break
		} else if sta.IsEmpty() && !stb.IsEmpty() { // a 空了
			if flag {
				sta.Push("1")
				flag = false
			} else {
				for !stb.IsEmpty() {
					result.Push(stb.Pop())
				}
				break
			}
		} else if !sta.IsEmpty() && stb.IsEmpty() { // b 空了
			if flag {
				stb.Push("1")
				flag = false
			} else {
				for !sta.IsEmpty() {
					result.Push(sta.Pop())
				}
				break
			}
		}
	}

	s := ""
	for !result.IsEmpty() {
		s = s + result.Pop()
	}

	return s
}

type Stack struct {
	Elem []string
}

func ConstructStack() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	if len(s.Elem) == 0 {
		return true
	}
	return false
}

func (s *Stack) Size() int {
	return len(s.Elem)
}

func (s *Stack) Push(e string) {
	s.Elem = append(s.Elem, e)
}

func (s *Stack) Pop() (e string) {
	if s.IsEmpty() {
		return ""
	}
	e = s.Elem[len(s.Elem)-1]
	s.Elem = s.Elem[0 : len(s.Elem)-1]
	return e
}

func (s *Stack) Top() (e string) {
	if !s.IsEmpty() {
		return s.Elem[len(s.Elem)-1]
	}
	return ""
}
