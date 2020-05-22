/**
 * @Time : 2020-05-14 17:41
 * @Author : yz
 */

package types

type Stack struct {
	Elem []interface{}
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

func (s *Stack) Push(e interface{})  {
	s.Elem = append(s.Elem, e)
}

func (s *Stack) Pop() (e interface{}) {
	if s.IsEmpty() {
		return nil
	}
	e = s.Elem[len(s.Elem)-1]
	s.Elem = s.Elem[0:len(s.Elem)-1]
	return e
}

func (s *Stack) Top() (e interface{}) {
	if !s.IsEmpty() {
		return s.Elem[len(s.Elem)-1]
	}
	return nil
}

