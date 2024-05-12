package brainteaser

import (
	"fmt"
	"testing"
)

func Test_nthPersonGetsNthSeat(t *testing.T) {
	//fmt.Println(nthPersonGetsNthSeat(3))
	s := []int{1, 3, 5, 7}
	for i, v := range s {
		if v == 7 {
			s = append(s[:i], s[i+1:]...)
		}
	}
	fmt.Println(s)
}
