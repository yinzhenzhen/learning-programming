package heap

import (
	"fmt"
	"testing"
)

func Test_findKthLargest(t *testing.T) {
	//fmt.Println(findKthLargest([]int{3,2,3,1,2,4,5,5,6}, 4))
	fmt.Println(findKthLargest([]int{4, 6, 8, 5, 9, 7}, 4))
	//fmt.Println(findKthLargest([]int{50,10,90,30,70,40,80,60,20}, 2))
}

func Test_nthUglyNumber(t *testing.T) {
	fmt.Println(nthUglyNumber(95))
}
