/**
 * @Time : 2020-12-01 14:09
 * @Author : yz
 */

package swordfinger

import (
	"strings"
)

/**
替换空格
*/

func replaceSpace(s string) string {
	temp := strings.Split(s, " ")
	return strings.Join(temp, "%20")
}
