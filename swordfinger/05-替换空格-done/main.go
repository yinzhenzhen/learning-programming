/**
 * @Time : 2020-12-01 14:09
 * @Author : yz
 */

package main

import (
	"fmt"
	"strings"
)

func replaceSpace(s string) string {
	temp := strings.Split(s, " ")
	return strings.Join(temp, "%20")
}

func main() {
	fmt.Println(replaceSpace("We are happy."))
}
