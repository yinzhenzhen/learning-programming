/**
 * @Time : 2020-09-11 19:36
 * @Author : yz
 */

package huawei

func wordCount(str string, single string) (count int) {
	for _, v := range str {
		if string(v) == single || string(v-32) == single || string(v+32) == single {
			count++
		}
	}
	return
}
