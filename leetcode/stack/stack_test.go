package stack

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/yinzhenzhen/learning-programming/types"
	"testing"
)

func Test_isValidParentheses(t *testing.T) {
	b := isValidParentheses("[]{")
	require.Equal(t, false, b)
	b = isValidParentheses("()[]{}")
	require.Equal(t, true, b)
}

func Test_trap(t *testing.T) {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func Test_simplifyPath(t *testing.T) {
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
}

func Test_evalRPN(t *testing.T) {
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

func Test_minStack(t *testing.T) {
	minStack := NewMinStack()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}

func Test_BST(t *testing.T) {
	null := -1
	elem := []int{7, 3, 15, null, null, 9, 20}
	root := &types.TreeNode{elem[0], nil, nil}
	root.ConstructTree(elem, 1, 2)
	bst := NewBSTIterator(root)
	fmt.Println(bst.elems)
	fmt.Println(bst.Next())
	fmt.Println(bst.Next())
	fmt.Println(bst.HasNext())
	fmt.Println(bst.Next())
	fmt.Println(bst.Next())
	fmt.Println(bst.Next())
	fmt.Println(bst.Next())
	fmt.Println(bst.Next())
	fmt.Println(bst.HasNext())
	fmt.Println(bst.Next())
}

func Test_calculate(t *testing.T) {
	fmt.Println(calculate("2-(5-6)"))
	//fmt.Println(combine([]int{1, 2, 3}))
}

func Test_backspaceCompare(t *testing.T) {
	fmt.Println(backspaceCompare("abcd", "bbcd"))
}

func Test_decodeAtIndex(t *testing.T) {
	fmt.Println(decodeAtIndex("y959q969u3hb22odq595", 222280369))
}

func Test_StockSpanner(t *testing.T) {
	obj := NewStockSpanner()
	//fmt.Println(obj.Next(100))
	//fmt.Println(obj.Next(80))
	//fmt.Println(obj.Next(60))
	//fmt.Println(obj.Next(70))
	//fmt.Println(obj.Next(60))
	//fmt.Println(obj.Next(75))
	//fmt.Println(obj.Next(85))
	fmt.Println(obj.Next(29))
	fmt.Println(obj.Next(91))
	fmt.Println(obj.Next(62))
	fmt.Println(obj.Next(76))
	fmt.Println(obj.Next(51))
}

func Test_minAddToMakeValid(t *testing.T) {
	fmt.Println(minAddToMakeValid("()"))
}

func Test_validateStackSequences(t *testing.T) {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}

func Test_isValid(t *testing.T) {
	fmt.Println(isValid("aabcbc"))
}

func Test_nextLargerNodes(t *testing.T) {
	head := types.ConstructLink([]int{1, 2, 3})
	head.Print()
}

func Test_removeOuterParentheses(t *testing.T) {
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
	//s := stack.Construct()
	//s.Push(1)
	//s.Push(2)
	//s.Push(3)
	//fmt.Println(s.Pop())
	//fmt.Println(s.Pop())
	//fmt.Println(s.Pop())
}

func Test_removeDuplicates(t *testing.T) {
	fmt.Println(removeDuplicates("abbcc"))
}
