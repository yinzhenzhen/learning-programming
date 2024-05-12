package dynamic_programming

import (
	"fmt"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("a"))
}

func Test_uniquePathsWithObstacles(t *testing.T) {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 1}}))

	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0},
		{1, 1},
		{0, 0},
	}))

	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}))
}

func Test_minPathSum(t *testing.T) {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}

func Test_numDecodings(t *testing.T) {
	fmt.Println(numDecodings("10011"))
}
