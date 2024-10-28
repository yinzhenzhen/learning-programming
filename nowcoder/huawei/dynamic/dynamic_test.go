package dynamic

import (
	"fmt"
	"testing"
)

func Test_levenshtein(t *testing.T) {
	//levenshtein()
	c := lev("abcdefg", "abcdef")
	fmt.Println(c)
}
