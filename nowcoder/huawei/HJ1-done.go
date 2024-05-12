/**
 * @Time : 2020-09-11 19:14
 * @Author : yz
 */

package huawei

import (
	"strings"
)

func lastWordLength(str string) int {
	if str == "" {
		return 0
	}
	words := strings.Split(str, " ")
	return len(words[len(words)-1])
}
