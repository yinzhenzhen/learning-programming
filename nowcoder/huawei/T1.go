/**
 * @Time : 2020-09-15 19:10
 * @Author : yz
 */

package huawei

import (
	"strings"
)

func minStr(str string) (result string) {

	// 用来进行对比
	data := str

	// 记录之前试图能替换其他元素的元素
	// 被替换元素的位置
	var indexA = -1
	// 替换元素的位置
	var indexB = -1

	for i := 1; i < len(str); i++ {
		for j := 0; j < i; j++ {
			// 如果元数据的当前元素比data小，记录
			if str[i] < data[j] {

				// 如果之前存在期望替换的元素
				if indexA != -1 && indexB != -1 {
					// 判断新的元素是否比之前的元素还小
					if str[i] < str[indexB] {
						indexB = i
					}
				} else {
					indexA = j
					indexB = i
					break
				}
			}
		}
	}

	// 需要进行元素替换
	if indexA != -1 && indexB != -1 {
		temp := strings.Split(str, "")
		temp[indexA] = string(str[indexB])
		temp[indexB] = string(str[indexA])
		result = strings.Join(temp, "")
	} else {
		result = str
	}

	return
}
