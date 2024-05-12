/**
 * @Time : 2020-05-12 17:00
 * @Author : yz
 */

package stack

type MinStack struct {
	Elem []int
}

/** initialize your data structure here. */
func NewMinStack() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.Elem = append(this.Elem, x)
}

func (this *MinStack) Pop() {
	this.Elem = this.Elem[0 : len(this.Elem)-1]
}

func (this *MinStack) Top() int {
	return this.Elem[len(this.Elem)-1]
}

func (this *MinStack) GetMin() int {
	var min = this.Elem[0]
	for i := 1; i < len(this.Elem); i++ {
		if this.Elem[i] < min {
			min = this.Elem[i]
		}
	}
	return min
}
