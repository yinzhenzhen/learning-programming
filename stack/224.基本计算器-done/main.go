/**
 * @Time : 2020-05-22 13:37
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-leetcode/types"
	"strconv"
)

func calculate(s string) int {
	dataStack := types.ConstructStack()
	opStack := types.ConstructStack()

	datas := []int{}

	for i := 0; i < len(s); i++ {
		str := fmt.Sprintf("%c", s[i])

		if str == "+" || str == "-" {
			opStack.Push(str)
			fmt.Println("op push ", str)
		} else if str == " " {
			continue
		} else if str == "(" {
			opStack.Push(str)
			fmt.Println("op push ", str)
		} else if str == ")" {
			op := opStack.Pop().(string)
			fmt.Println("op pop ", op)
			// 如果取出来的第一个操作符是（，说明是(num)的情况
			if op == "(" {
				if opStack.Size() > 0 {
					op = opStack.Pop().(string)
					fmt.Println("op pop ", op)
				} else {
					continue
				}
			}
			// 取出两个数据, 一个操作符, 取出操作符的"("
			stackData1 := dataStack.Pop().(int)
			fmt.Println("data pop ", str)
			stackData2 := dataStack.Pop().(int)
			fmt.Println("data pop ", str)
			if op == "+" {
				stackData2 = stackData2 + stackData1
			} else if op == "-" {
				stackData2 = stackData2 - stackData1
			}
			dataStack.Push(stackData2)
			fmt.Println("data push ", str)

		} else {
			newData, _ := strconv.Atoi(str)

			// 如果这是最后一个元素，
			if i == len(s)-1 {
				datas = append(datas, newData)
				newData = combine(datas)
			} else {
				nextStr := fmt.Sprintf("%c", s[i+1])

				// 如果下一个不是数字，说明目前datas的数字可以组合
				if nextStr == "+" || nextStr == "-" || nextStr == "(" || nextStr == ")" || nextStr == " " {
					// 将连续数字组合
					datas = append(datas, newData)
					newData = combine(datas)
					datas = []int{}
				//  下一个还是数字
				} else {
					datas = append(datas, newData)
					continue
				}
			}

			if opStack.Size() > 0 {
				op := opStack.Top().(string)
				if op != "+" && op != "-" {
					dataStack.Push(newData)
					fmt.Println("data push ", newData)
					continue
				} else {
					stackData := dataStack.Pop().(int)
					fmt.Println("data pop ", stackData)
					opStack.Pop()
					fmt.Println("op pop ", op)
					if op == "+" {
						stackData = stackData + newData
					} else if op == "-" {
						stackData = stackData - newData
					}
					dataStack.Push(stackData)
					fmt.Println("data push ", stackData)
				}
			} else {
				dataStack.Push(newData)
				fmt.Println("data push ", newData)
			}


		}
	}

	//var result = dataStack.Elem[0].(int)
	//
	//for i,j := 1,0; i < len(dataStack.Elem); i,j = i+1,j+1 {
	//	stackData := dataStack.Elem[i].(int)
	//	op := opStack.Elem[j].(string)
	//	if op == "+" {
	//		result = result + stackData
	//	} else if op == "-" {
	//		result = result - stackData
	//	}
	//}

	return dataStack.Top().(int)
}

func combine(datas []int) (result int) {
	j := 1
	for i := len(datas) - 1; i >= 0; i-- {
		result += datas[i] * j
		j = j * 10
	}
	return result
}

func main()  {
	fmt.Println(calculate("2-(5-6)"))
	//fmt.Println(combine([]int{1, 2, 3}))
}