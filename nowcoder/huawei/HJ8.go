/**
 * @Time : 2020-09-14 23:24
 * @Author : yz
 */

package huawei

import (
	"fmt"
	"sort"
)

func main() {

	var num int

	for {
		m := make(map[int]int)
		s := []int{}
		n, _ := fmt.Scan(&num)
		if n == 0 {
			break
		}

		for i := 0; i < num; i++ {

			var tempIndex, tempData int
			_, _ = fmt.Scan(&tempIndex, &tempData)

			if _, ok := m[tempIndex]; ok {
				m[tempIndex] = m[tempIndex] + tempData
			} else {
				m[tempIndex] = tempData
				s = append(s, tempIndex)
			}
		}

		sort.Ints(s)

		for i := 0; i < len(s); i++ {
			fmt.Println(s[i], m[s[i]])
		}
	}

}
