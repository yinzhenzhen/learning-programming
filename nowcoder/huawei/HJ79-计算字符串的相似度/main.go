/**
 * @Time : 2020-09-13 18:14
 * @Author : yz
 */

package main

import "fmt"

func calculateStringDistance(strA, strB string) (result string) {



	return
}

func main()  {
	var strA, strB string
	for {
		n, _ := fmt.Scan(&strA, &strB)
		if n == 0 {
			break
		}

		fmt.Println(calculateStringDistance(strA, strB))
	}
}