/**
 * @Time : 2020-09-11 19:14
 * @Author : yz
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func lastWordLength(str string) int {
	if str == "" {
		return 0
	}

	words := strings.Split(str, " ")

	return len(words[len(words)-1])
}


func main()  {

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