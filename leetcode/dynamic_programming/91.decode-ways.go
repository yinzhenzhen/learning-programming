/**
 * @Time : 2020-12-27 15:59
 * @Author : yz
 */

package dynamic_programming

import (
	"strconv"
	"strings"
)

// f[i] = f[i-1] + f[i-2]  11
// f[i] = f[i-1]           01 / 27
// f[i] = f[i-2]           10
func numDecodings(s string) int {

	if len(s) == 1 {
		if s == "0" {
			return 0
		}
		return 1
	}

	length := len(s)
	f := make([]int, length)

	ss := strings.Split(s, "")

	if ss[0] != "0" {
		f[0] = 1
	} else {
		return 0
	}

	sss := ss[0] + ss[1]
	temp, _ := strconv.Atoi(sss)
	if ss[1] != "0" {
		if temp <= 26 {
			f[1] = 2
		} else {
			f[1] = 1
		}
	} else {
		if temp > 26 {
			return 0
		}
		f[1] = f[0]
	}

	for i := 2; i < len(s); i++ {

		sss := ss[i-1] + ss[i]
		temp, _ := strconv.Atoi(sss)

		if ss[i] == "0" {
			if ss[i-1] == "0" || temp > 26 {
				return 0
			}
			f[i] = f[i-2]

		} else if temp > 26 {
			f[i] = f[i-1]
		} else {
			if ss[i-1] == "0" {
				f[i] = f[i-1]
			} else {
				f[i] = f[i-1] + f[i-2]
			}

		}

	}

	return f[length-1]
}
