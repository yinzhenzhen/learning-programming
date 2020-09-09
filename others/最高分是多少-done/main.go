/**
 * @Time : 2020-09-08 09:38
 * @Author : yz
 */

package main

import (
	"fmt"
	"io"
	"sync"
)

type HighestScore struct {
	m sync.RWMutex
	// 根据成绩进行从高到低排序
	all []*Student
}

type Student struct {
	Id int
	Score int
}

// 对学生的成绩进行从高到低的排序
func (hs *HighestScore) Init() {

	// 选择排序
	for i := 0; i < len(hs.all); i++ {
		high := i
		for j := i+1; j < len(hs.all); j++ {
			if hs.all[j].Score > hs.all[high].Score {
				high = j
			}
		}
		if high != i {
			temp := hs.all[high]
			hs.all[high] = hs.all[i]
			hs.all[i] = temp
		}
	}

}

// 查询指定id之间分数最高的
func (hs *HighestScore) QueryScore(start, end int) int {
	//fmt.Println("enter query")

	// 遍历所有学生，最先属于start-end之间的就是这个区间的最高分
	for _, stu := range hs.all {
		if (stu.Id >= start && stu.Id <= end) || (stu.Id <= start && stu.Id >= end) {
			return stu.Score
		}
	}

	return 0
}

// 更新成绩
func (hs *HighestScore) UpdateScore(stuId, stuScore int)  {
	//fmt.Println("enter update")

	index := 0

	forward := false
	back := false

	//fmt.Println(hs.all)

	for i, stu := range hs.all {
		if stu.Id == stuId {

			// 新分数比当前分数高
			if stu.Score < stuScore  {
				index = i
				forward = true

			// 新分数比当前分数低
			} else if stu.Score > stuScore {
				index = i
				back = true
			}

			stu.Score = stuScore
			//fmt.Println(stu.Score)
			break
		}
	}

	//fmt.Println(forward, back)
	//fmt.Println(hs.all)

	var news []*Student

	// 如果新分数比当前分数高，往前移动
	if forward && !back {
		flag := false
		for i := 0; i < index; i++ {
			//fmt.Println(hs.all[index].Score, hs.all[i].Score)

			// 如果更新后的成绩比当前学生的成绩高
			if hs.all[index].Score > hs.all[i].Score {

				news = append(news, hs.all[0:i]...)
				news = append(news, hs.all[index])
				news = append(news, hs.all[i:index]...)
				news = append(news, hs.all[index+1:]...)
				flag = true
				break
			}
		}
		// 发生变化
		if flag {
			hs.all = news
		}

	// 如果新分数比当前分数低，往后移动
	} else if !forward && back {
		flag := false
		for i := index+1; i < len(hs.all); i++ {

			if hs.all[index].Score > hs.all[i].Score {

				news = append(news, hs.all[0:index]...)
				news = append(news, hs.all[index+1:i]...)
				news = append(news, hs.all[index])
				news = append(news, hs.all[i:]...)
				flag = true
				break
			}
		}
		// 发生变化
		if flag {
			hs.all = news
		}

	} else {
		// 分数没有发生改变
		return
	}

	//fmt.Println(hs.all)
}

type Operator struct {
	op string
	a int
	b int
}

func main() {

	for ;  ; {
		var n, m int
		var ss []*Student

		_, err := fmt.Scan(&n, &m)
		if err == io.EOF {
			break
		}

		for i := 1; i <= n; i++ {
			var score int
			_, _ = fmt.Scan(&score)
			s := &Student{
				i,
				score,
			}
			ss = append(ss, s)
		}

		hs := &HighestScore{
			all: ss,
		}

		hs.Init()

		var operator []Operator
		for i := 1; i <= m; i++  {
			var op string
			var a, b int
			_, _ = fmt.Scan(&op, &a, &b)
			temp := Operator{
				op,
				a,
				b,
			}
			operator = append(operator, temp)
		}

		for i := 0; i < len(operator); i++  {
			if operator[i].op == "Q" {
				res := hs.QueryScore(operator[i].a, operator[i].b)
				fmt.Println(res)
			} else if operator[i].op == "U" {
				hs.UpdateScore(operator[i].a, operator[i].b)
			} else {
				fmt.Println("error op")
			}

		}
	}


}
