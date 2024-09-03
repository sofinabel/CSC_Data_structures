package main

import (
	"bufio"
	"fmt"
	"os"
)

type minHeap struct {
	items     []int
	size      int
	swaps     [][]int
	swapCount int
}

func newMinHeap(arr []int, n int) *minHeap {
	return &minHeap{items: arr, size: n, swaps: make([][]int, 4*n), swapCount: 0}
}

func (h *minHeap) buildMinHeap() {
	for i := h.size/2 - 1; i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *minHeap) siftDown(i int) {
	minInd := i
	l := getLeftChild(i)
	if l < h.size && h.items[l] < h.items[minInd] {
		minInd = l
	}
	r := getRightChild(i)
	if r < h.size && h.items[r] < h.items[minInd] {
		minInd = r
	}
	if i != minInd {
		h.swaps[h.swapCount] = make([]int, 2)
		h.swaps[h.swapCount][0], h.swaps[h.swapCount][1] = i, minInd
		h.swapCount++
		//fmt.Println(i, minInd)
		swap(&h.items[i], &h.items[minInd])
		h.siftDown(minInd)
	}
}

func getRightChild(i int) int {
	return 2*i + 1
}

func getLeftChild(i int) int {
	return 2*i + 2
}

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	heap := newMinHeap(arr[:], n)
	heap.buildMinHeap()
	fmt.Fprintln(out, heap.swapCount)
	for i := 0; i < heap.swapCount; i++ {
		fmt.Fprintln(out, heap.swaps[i][0], heap.swaps[i][1])
	}
}
