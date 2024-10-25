package huawei

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func Test_lastWordLength(t *testing.T) {
	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		if input == "" {
			return
		}
		input = input[:len(input)-1]
		fmt.Println(lastWordLength(input))
	}
}

func Test_wordCount(t *testing.T) {
	var str string
	var single string
	for {
		n, _ := fmt.Scan(&str, &single)
		if n == 0 {
			return
		}
		fmt.Println(wordCount(str, single))
	}
}

func Test_random(t *testing.T) {
	var num int

	for {
		var input []int
		n, _ := fmt.Scan(&num)
		if n == 0 {
			return
		}

		for i := 0; i < num; i++ {
			var temp int
			_, _ = fmt.Scan(&temp)
			input = append(input, temp)
		}

		output := random(input)
		for _, v := range output {
			fmt.Println(v)
		}

	}
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

func Test_H25(t *testing.T) {
	H25()
}

func Test_HJ51(t *testing.T) {
	HJ51()
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

func Test_OR101(t *testing.T) {
	OR101()
}

func Test_minStr(t *testing.T) {
	var str string

	for {
		n, _ := fmt.Scan(&str)
		if n == 0 {
			break
		}
		fmt.Println(minStr(str))
	}
	//fmt.Println(minStr("bcdefa"))
}

func Test_T2(t *testing.T) {
	T2()
}

func Test_putApples(t *testing.T) {
	putApples()
}

func Test_leastCommonMultiple(t *testing.T) {
	leastCommonMultiple()
}

func Test_levenshtein(t *testing.T) {
	levenshtein()
}
