/**
 * @Time : 2020-06-03 09:46
 * @Author : yz
 */

package sort

type Weight struct {
	num    int
	weight int
}

func getKth(lo int, hi int, k int) int {
	// key为数字，value为对应权重
	weights := []Weight{}
	for i := lo; i <= hi; i++ {
		w := Weight{num: i}
		w.calculateWeight()
		weights = append(weights, w)
	}

	// 根据权重进行排序
	// 冒泡排序
	flag := true
	for i := 0; i < len(weights) && flag; i++ {
		flag = false
		for j := len(weights) - 2; j >= i; j-- {
			word1 := weights[j]
			word2 := weights[j+1]
			if word1.weight > word2.weight {
				swap3(weights, j, j+1)
				flag = true
			}
		}
	}
	return weights[k-1].num
}

// 交换
func swap3(word []Weight, i int, j int) {
	temp := word[i]
	word[i] = word[j]
	word[j] = temp
}

func (w *Weight) calculateWeight() {
	data := w.num
	count := 0
	for data != 1 {
		// 为偶数
		if data%2 == 0 {
			data = data / 2
		} else {
			data = 3*data + 1
		}
		count++
	}
	w.weight = count
	return
}
