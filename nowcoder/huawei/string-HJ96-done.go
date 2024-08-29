/**
 * @Time : 2020-09-13 17:05
 * @Author : yz
 */

package huawei

const mark string = "*"

func addMark(data string) (result string) {
	// 数字是否连续
	flag := false

	// 暂时存放连续数字的字符串
	var temp string

	// 遍历处理每个字符
	for i := 0; i < len(data); i++ {

		// 如果当前字符为数字
		if data[i] >= 48 && data[i] <= 57 {

			if flag == true {
				// 如果前一个字符也是数字，将它放到连续数字切片中
				temp = temp + string(data[i])

			} else {
				// 前一个字符不是数字，但是当前字符是数字
				temp = temp + string(data[i])
				flag = true
			}

			// 如果当前字符不为数字
		} else {
			// 如果当前字符不是数字，前一个字符为数字，将暂存数字添加到结果中，并将标志位置为false
			if flag == true {

				result = result + mark + temp + mark + string(data[i])

				temp = ""
				flag = false

				// 前一个字符不为数字，当前也不为数字，直接放到结果中
			} else {
				result = result + string(data[i])
			}
		}
	}

	if temp != "" {
		result = result + mark + temp + mark
	}

	return
}
