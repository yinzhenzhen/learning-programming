/**
 * @Time : 2020-09-13 23:53
 * @Author : yz
 */

package huawei

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func combine(all []Interval) (result []Interval) {

	result = append(result, all[0])

	for i := 1; i < len(all); i++ {
		c1 := result[len(result)-1]
		c2 := all[i]

		if c1.start <= c2.start && c1.end >= c2.end {

			temp := Interval{c1.start, c1.end}
			result = append(result[:len(result)-1], temp)

		} else if c1.start >= c2.start && c1.end <= c2.end {

			temp := Interval{c2.start, c2.end}
			result = append(result[:len(result)-1], temp)

		} else if c1.start <= c2.start && c1.end >= c2.start {

			temp := Interval{c1.start, c2.end}
			result = append(result[:len(result)-1], temp)

		} else if c1.start >= c2.start && c1.start <= c2.end {

			temp := Interval{c2.start, c1.end}
			result = append(result[:len(result)-1], temp)

		} else {

			result = append(result, c2)

		}
	}

	return
}

type Interval struct {
	start, end int
}

func OR101() {
	var str string

	for {
		inputReader := bufio.NewReader(os.Stdin)

		str, _ = inputReader.ReadString('\n')

		if str == "" {
			return
		}

		strSlice := strings.Split(str[:len(str)-1], " ")

		var all []Interval

		for _, v := range strSlice {
			data := strings.Split(v, ",")

			if len(data) != 2 {
				break
			}

			temp1, _ := strconv.Atoi(data[0])
			temp2, _ := strconv.Atoi(data[1])

			inter := Interval{temp1, temp2}
			all = append(all, inter)
		}

		for i := 0; i < len(all); i++ {
			min := i

			for j := i + 1; j < len(all); j++ {
				if all[j].start < all[min].start {
					min = j
				}
			}

			if min != i {
				temp := all[i]
				all[i] = all[min]
				all[min] = temp
			}
		}

		result := combine(all)

		ss := ""

		for i, v := range result {
			s := fmt.Sprintf("%d,%d", v.start, v.end)
			ss += s
			if i != len(result) {
				ss += " "
			}
		}

		fmt.Println(ss)
	}
}
