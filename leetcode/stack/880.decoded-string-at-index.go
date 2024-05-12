/**
 * @Time : 2020-05-26 13:20
 * @Author : yz
 */

package stack

import (
	"fmt"
	"strconv"
)

func decodeAtIndex(S string, K int) string {

	var newS string

	for i := 0; i < len(S); i++ {
		str := fmt.Sprintf("%c", S[i])
		// 如果是字母，写入newS
		if str >= "a" && str <= "z" {
			newS = newS + str
			// 如果是数字，进行重复
		} else {
			count, _ := strconv.Atoi(str)
			temp := newS
			for count > 1 {
				// 运行环境，panic
				newS = newS + temp
				count--
			}

			if len(newS) >= K {
				//fmt.Println(newS)
				return fmt.Sprintf("%c", newS[K-1])
			}
		}
	}

	//fmt.Println(newS)

	return fmt.Sprintf("%c", newS[K-1])
}
