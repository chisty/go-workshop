package main

import (
	"fmt"
)

func main() {
	fmt.Println("Max Heap")

	testMaxHeap([]int{2, 7, 14, 11, 8, 10})
	testMaxHeap([]int{1, 2, 3, 4, 5, 6, 7})
	testMaxHeap([]int{100, 17, 19, 23, 50, 17})
}

func testMaxHeap(input []int) {
	fmt.Println("Input= ", input)
	heap := newMaxHeap(len(input))

	for _, value := range input {
		heap.insert(value)
	}
	fmt.Println("Result= ", heap.heap)
	fmt.Println("DownHeapify=")

	for i := 0; i < len(input); i++ {
		res := heap.remove()
		// fmt.Println("Max= ", res, "Heap= ", heap.heap)
		fmt.Printf("%d\t", res)
	}

	fmt.Println("\n-----------------------------------------------------------------------------")
}

func (m *maxHeap) insert(item int) error {
	if m.isFull() {
		return fmt.Errorf("Cannot insert. Heap is full")
	}

	m.heap = append(m.heap, item)
	m.upHeapify(m.size)
	m.size++
	return nil
}

func (m *maxHeap) remove() int {
	if m.isEmpty() {
		return -1
	}
	res := m.heap[0]
	m.heap[0] = m.heap[m.size-1]
	m.heap = m.heap[:m.size-1]
	m.size--

	m.downHeapify(0)
	return res
}

func (m *maxHeap) downHeapify(index int) {
	left := leftIndex(index)
	right := rightIndex(index)

	if left > m.size-1 && right > m.size-1 {
		return
	}

	largest := index
	if left <= m.size-1 && m.heap[left] > m.heap[largest] {
		largest = left
	}
	if right <= m.size-1 && m.heap[right] > m.heap[largest] {
		largest = right
	}
	if largest != index {
		m.swap(largest, index)
		m.downHeapify(largest)
	}
}

func (m *maxHeap) upHeapify(index int) {
	for {
		parentIndex := parentIndex(index)
		if parentIndex < 0 {
			break
		}
		if m.heap[parentIndex] >= m.heap[index] {
			break
		}
		m.swap(index, parentIndex)
		index = parentIndex
	}
}

func parentIndex(index int) int {
	return (index - 1) / 2
}

func leftIndex(index int) int {
	return index*2 + 1
}

func rightIndex(index int) int {
	return index*2 + 2
}

func (m *maxHeap) swap(index, parentIndex int) {
	m.heap[index], m.heap[parentIndex] = m.heap[parentIndex], m.heap[index]
}

func (m *maxHeap) isFull() bool {
	if m.size == m.maxSize {
		return true
	}
	return false
}

func (m *maxHeap) isEmpty() bool {
	if m.size == 0 {
		return true
	}
	return false
}

func newMaxHeap(m int) *maxHeap {
	return &maxHeap{
		heap:    []int{},
		size:    0,
		maxSize: m,
	}
}

type maxHeap struct {
	heap    []int
	size    int
	maxSize int
}
