package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Processor struct {
	releaseTime int //приоритет в очереди
	procIndex   int
}

type PriorityQueue []*Processor

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].releaseTime == pq[j].releaseTime {
		return pq[i].procIndex < pq[j].procIndex
	}
	return pq[i].releaseTime < pq[j].releaseTime
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	proc := x.(*Processor)
	*pq = append(*pq, proc)
}

func (pq *PriorityQueue) Pop() interface{} {
	queue := *pq
	proc := queue[len(queue)-1]
	*pq = queue[:len(queue)-1]
	return proc
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, m int
	fmt.Fscan(in, &n, &m)
	duration := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &duration[i])
	}
	pq := make(PriorityQueue, 0)
	for i := 0; i < n; i++ {
		heap.Push(&pq, &Processor{releaseTime: 0, procIndex: i})
	}
	for i := 0; i < m; i++ {
		x := heap.Pop(&pq)
		proc := x.(*Processor)
		fmt.Fprintln(out, proc.procIndex, proc.releaseTime)
		proc.releaseTime += duration[i]
		heap.Push(&pq, proc)
	}
}
