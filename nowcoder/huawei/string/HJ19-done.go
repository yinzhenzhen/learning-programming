/**
 * @Time : 2020-09-08 17:19
 * @Author : yz
 */

package string

import (
	"fmt"
	"strconv"
	"strings"
)

// 1.记录最多8条错误记录，对相同的错误记录(即文件名称和行号完全匹配)只记录一条，错误计数增加；(文件所在的目录不同，文件名和行号相同也要合并)
// 2.超过16个字符的文件名称，只记录文件的最后有效16个字符；(如果文件名不同，而只是文件名的后16个字符和行号相同，也不要合并)
// 3.输入的文件可能带路径，记录文件名称不能带路径

type ErrorRecords struct {
	// key=文件名称+行号，value为出现次数
	recordsMap map[string]*Description
	// 初次统计的结果
	recordSlice []Description
}

type Description struct {
	// 文件名称
	fileName string
	// 行号
	lineNumber int
	// 出现次数
	count int
	// 出现顺序，1为最先
	order int
}

func (er *ErrorRecords) Integration() []Description {
	var result []Description

	// 查找出现次数最高的八条记录
	for i := 0; i < len(er.recordSlice); i++ {

		max := i

		for j := i + 1; j < len(er.recordSlice); j++ {

			// 如果新的数据比最大的更先出现，标记为最大
			if er.recordSlice[j].order < er.recordSlice[max].order {
				max = j
			}

			//// 如果新的数据比最大的大，标记为最大
			//if er.recordSlice[j].count > er.recordSlice[max].count {
			//	max = j
			//
			//// 	如果新的数据和最大相等
			//} else if er.recordSlice[j].count == er.recordSlice[max].count {
			//
			//	// 如果新的数据比最大的更先出现，标记为最大
			//	if er.recordSlice[j].order < er.recordSlice[max].order {
			//		max = j
			//	}
			//
			//}
		}

		result = append(result, er.recordSlice[max])

		if max != i {
			temp := er.recordSlice[i]
			er.recordSlice[i] = er.recordSlice[max]
			er.recordSlice[max] = temp
		}

	}

	if len(result) > 8 {
		return result[len(result)-8:]
	}

	return result
}

func H19() {

	var fp string
	var num int

	ers := &ErrorRecords{
		recordsMap: make(map[string]*Description),
	}

	for {
		n, _ := fmt.Scanln(&fp, &num)
		if n == 0 {
			break
		}

		desp := &Description{
			lineNumber: num,
			count:      1,
		}

		if len(strings.Split(fp, "\\")) > 1 {
			temp := strings.Split(fp, "\\")
			desp.fileName = temp[len(temp)-1]
		} else {
			desp.fileName = fp
		}

		key := desp.fileName + "_" + strconv.Itoa(desp.lineNumber)

		if _, ok := ers.recordsMap[key]; ok {
			// 已经存在
			ers.recordsMap[key].count++
		} else {
			// 出现顺序
			desp.order = len(ers.recordsMap) + 1
			// 新元素
			ers.recordsMap[key] = desp
		}
	}

	for _, v := range ers.recordsMap {
		ers.recordSlice = append(ers.recordSlice, *v)
	}

	results := ers.Integration()

	for _, r := range results {
		if len(r.fileName) > 16 {
			fmt.Println(r.fileName[len(r.fileName)-16:], r.lineNumber, r.count)
		} else {
			fmt.Println(r.fileName, r.lineNumber, r.count)
		}

	}
}
