package string

import (
	"fmt"
	"testing"
)

func Test_lastWordLength(t *testing.T) {
	HJ1()
}

func Test_wordCount(t *testing.T) {
	//HJ2()
	fmt.Println(wordCount2("8 8 8  8A i i OOI              IIIaa", "A"))
}

func Test_split(t *testing.T) {
	var input1, input2 string

	_, _ = fmt.Scan(&input1, &input2)

	output1 := split(input1)
	output2 := split(input2)

	for _, v := range output1 {
		fmt.Println(v)
	}

	for _, v := range output2 {
		fmt.Println(v)
	}
}

func Test_baseConversion(t *testing.T) {
	var data string
	for {
		n, _ := fmt.Scan(&data)
		if n == 0 {
			break
		}
		fmt.Println(baseConversion(data[2:]))
	}
}

func Test_H19(t *testing.T) {
	H19()
}

func Test_addMark(t *testing.T) {
	var data string
	for {
		n, _ := fmt.Scan(&data)
		if n == 0 {
			break
		}
		fmt.Println(addMark(data))
	}
}

func Test_H106(t *testing.T) {
	//H106()
	fmt.Println(characterReverseOrder("I am a student"))
}

//func Test_OR101(t *testing.T) {
//	huawei.OR101()
//}
//
//func Test_minStr(t *testing.T) {
//	var str string
//
//	for {
//		n, _ := fmt.Scan(&str)
//		if n == 0 {
//			break
//		}
//		fmt.Println(huawei.minStr(str))
//	}
//	//fmt.Println(minStr("bcdefa"))
//}
//
//func Test_T2(t *testing.T) {
//	huawei.T2()
//}
