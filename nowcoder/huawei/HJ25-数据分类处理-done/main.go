/**
* @Time : 2020-09-13 19:03
* @Author : yz
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type IAndR struct {
	r string
	count string
	data  []string
}

func dataClassificationProcessing2(I, R []string) (result []string) {
	m := make(map[string]interface{})

	var iAndRs []*IAndR
	for i := 0; i < len(R); i++ {

		// 重复的r不用处理
		if _, ok := m[R[i]]; ok {
			continue
		}

		// 标记重复的r
		m[R[i]] = nil

		iAndRs = append(iAndRs, &IAndR{r: R[i]})
	}

	if len(iAndRs) == 0 {
		return nil
	}

	var wg sync.WaitGroup
	wg.Add(len(iAndRs))

	myFunc := func(iAndR *IAndR) {
		defer wg.Done()

		// 遍历I的元素
		for j := 0; j < len(I); j++  {

			f := strings.Contains(I[j], iAndR.r)

			// 如果当前i满足r的包含
			if f {
				// 第一个满足r的i
				iAndR.data = append(iAndR.data, strconv.Itoa(j))
				iAndR.data = append(iAndR.data, I[j])
			}
		}

		iAndR.count = strconv.Itoa(len(iAndR.data)/2)
	}

	for i := 0; i < len(iAndRs); i++ {
		go myFunc(iAndRs[i])
	}

	wg.Wait()

	temp := []string{}

	for _, iandr := range iAndRs {
		if iandr.count == "0" {
			continue
		}
		temp = append(append(append(temp, iandr.r), iandr.count), iandr.data...)
	}

	result = append(append(result, strconv.Itoa(len(temp))), temp...)

	return
}


func dataClassificationProcessing(I, R []string) (result []string) {
	m := make(map[string]interface{})

	for i := 0; i < len(R); i++ {

		// 重复的r不用处理
		if _, ok := m[R[i]]; ok {
			continue
		}

		// 标记重复的r
		m[R[i]] = nil

		//
		temp := []string{}

		// 遍历I的元素
		for j := 0; j < len(I); j++  {

			f := strings.Contains(I[j], R[i])

			// 如果当前i满足r的包含
			if f {
				// 第一个满足r的i
				temp = append(temp, strconv.Itoa(j))
				temp = append(temp, I[j])
			}
		}

		if len(temp) > 0 {
			result = append(result, R[i])
			result = append(result, strconv.Itoa(len(temp)/2))
			result = append(result, temp...)
		}

	}

	size := len(result)
	temp := []string{strconv.Itoa(size)}
	temp = append(temp, result...)
	result = temp

	return
}

func main()  {
	var INum, RNum int

	for {
		var I, R []string
		var RInt []int

		n, _ := fmt.Scan(&INum)

		if n == 0 {
			break
		}

		for i := 0; i < INum; i++ {
			var temp string
			_, _ = fmt.Scan(&temp)
			I = append(I, temp)
		}

		_, _ = fmt.Scan(&RNum)

		for i := 0; i < RNum; i++ {
			var temp int
			_, _ = fmt.Scan(&temp)
			RInt = append(RInt, temp)
		}

		sort.Ints(RInt)

		for _, v := range RInt {
			R = append(R, strconv.Itoa(v))
		}

		result := dataClassificationProcessing(I, R)

		for _, data := range result {
			fmt.Print(data + " ")
		}
		fmt.Println()
	}
}