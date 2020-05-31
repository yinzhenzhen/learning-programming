/**
 * @Time : 2020-05-30 21:58
 * @Author : yz
 */

package types

import "fmt"

type Queue struct {
	Elem []interface{}
	Length int
}

func ConstructQueue() *Queue {
	return &Queue{}
}

func (q *Queue) EnQueue(e interface{})  {
	q.Elem = append(q.Elem, e)
	q.Length++
}

func (q *Queue) DeQueue() (e interface{}) {
	if q.IsEmpty() {
		return nil
	}
	e = q.Elem[0]
	q.Elem = q.Elem[1:]
	q.Length--
	return
}

func (q *Queue) Size() int {
	return q.Length
}

func (q *Queue) IsEmpty() bool {
	if q.Length == 0 {
		return true
	}
	return false
}

func (q *Queue) Tail() (e interface{}) {
	if !q.IsEmpty() {
		return q.Elem[q.Length-1]
	}
	return nil
}

func (q *Queue) PrintQueue() {
	for i := 0; i < q.Length-1; i++ {
		switch q.Elem[i].(type) {
		case string:
			str := fmt.Sprintf("%s", q.Elem[i])
			fmt.Printf(str + "->")
		case int:
			str := fmt.Sprintf("%d", q.Elem[i])
			fmt.Printf(str + "->")
		case byte:
			str := q.Elem[i].(byte)
			fmt.Printf(string(str) + "->")
		}

	}
	fmt.Printf("%c\n", q.Elem[q.Length-1])
}