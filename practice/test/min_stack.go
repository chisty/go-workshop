package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Min Stack")

	minStack := Constructor()
	minStack.Push(3)
	minStack.Push(10)
	minStack.Push(5)
}

type MinStack struct {
	stack   []int
	minData []PairData
}

type PairData struct {
	value int
	pos   int
}

func NewPairData(v, p int) PairData {
	return PairData{value: v, pos: p}
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)

	i := sort.Search(len(this.minData), func(i int) bool {
		return this.minData[i].value <= x
	})

	this.minData = append(this.minData, NewPairData(0, 0))
	copy(this.minData[i+1:], this.minData[i:])
	this.minData[i] = NewPairData(x, len(this.stack))
}

func (this *MinStack) Pop() {
	pos := len(this.stack)

	this.stack = this.stack[:pos-1]

	delPos := 0

	for i := 0; i < len(this.minData); i++ {
		if this.minData[i].pos == pos {
			delPos = i
			break
		}
	}

	copy(this.minData[delPos:], this.minData[delPos+1:])
	this.minData = this.minData[:pos-1]
}

func (this *MinStack) Top() int {
	pos := len(this.stack)
	return this.stack[pos-1]
}

func (this *MinStack) GetMin() int {
	pos := len(this.minData)
	return this.minData[pos-1].value
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
