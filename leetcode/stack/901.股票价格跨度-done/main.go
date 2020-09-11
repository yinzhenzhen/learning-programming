/**
 * @Time : 2020-05-25 17:24
 * @Author : yz
 */

package main

import "fmt"

type StockSpanner struct {
	Elem []interface{}
}

var stock StockSpanner

func Constructor() StockSpanner {
	stock = StockSpanner{}
	return stock
}

func (s *StockSpanner) IsEmpty() bool {
	if len(s.Elem) == 0 {
		return true
	}
	return false
}

func (s *StockSpanner) Size() int {
	return len(s.Elem)
}

func (s *StockSpanner) Push(e interface{})  {
	s.Elem = append(s.Elem, e)
}

func (s *StockSpanner) Pop() (e interface{}) {
	if s.IsEmpty() {
		return nil
	}
	e = s.Elem[len(s.Elem)-1]
	s.Elem = s.Elem[0:len(s.Elem)-1]
	return e
}

func (s *StockSpanner) Top() (e interface{}) {
	if !s.IsEmpty() {
		return s.Elem[len(s.Elem)-1]
	}
	return nil
}

func (this *StockSpanner) Next(price int) int {

	// 首次输入
	if this.Top() == nil {
		this.Push(price)
		fmt.Println("this push ", price)
		return 1
	}

	i := this.Size() - 1
	result := 0

	for i >= 0 && this.Elem[i].(int) <= price {
		result++
		i--
	}

	this.Push(price)

	return result+1
}

func main()  {
	obj := Constructor()
	//fmt.Println(obj.Next(100))
	//fmt.Println(obj.Next(80))
	//fmt.Println(obj.Next(60))
	//fmt.Println(obj.Next(70))
	//fmt.Println(obj.Next(60))
	//fmt.Println(obj.Next(75))
	//fmt.Println(obj.Next(85))
	fmt.Println(obj.Next(29))
	fmt.Println(obj.Next(91))
	fmt.Println(obj.Next(62))
	fmt.Println(obj.Next(76))
	fmt.Println(obj.Next(51))
}