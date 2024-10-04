package main

import (
	"bufio"
	"fmt"
	"os"
)

type Elem struct {
	num    int
	curMax int
}

type Stack struct {
	items []Elem
}

func newStack() *Stack {
	return &Stack{items: []Elem{}}
}

func (s *Stack) Push(item int) {
	if s.isEmpty() {
		s.items = append(s.items, Elem{num: item, curMax: item})
	} else {
		last := s.Top()
		if last.curMax < item {
			s.items = append(s.items, Elem{num: item, curMax: item})
		} else {
			s.items = append(s.items, Elem{num: item, curMax: last.curMax})
		}
	}

}

func (s *Stack) Top() Elem {
	return s.items[s.Size()-1]
}

func (s *Stack) Pop() Elem {
	item := s.items[s.Size()-1]
	s.items = s.items[:s.Size()-1]
	return item

}

func (s *Stack) isEmpty() bool {
	if s.Size() == 0 {
		return true
	}
	return false
}

func (s *Stack) Size() int {
	return len(s.items)
}

type Queue struct {
	items []int
}

func NewQueue() *Queue {
	return &Queue{items: []int{}}
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, elem, m, maxx int
	queue := NewQueue()
	inStack := newStack()
	outStack := newStack()
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &elem)
		queue.Enqueue(elem)
	}

	fmt.Fscan(in, &m)

	for inStack.Size() < m {
		//fmt.Println("in")
		elem = queue.Dequeue()
		inStack.Push(elem)
	}

	for inStack.Size()+outStack.Size() == m {
		//массовая переброска
		if inStack.Size() == m && outStack.isEmpty() {
			for outStack.Size() < m {
				//fmt.Println("out")
				elem = inStack.Pop().num
				outStack.Push(elem)
			}
		}

		if inStack.isEmpty() {
			maxx = outStack.Top().curMax
		} else {
			maxx = max(inStack.Top().curMax, outStack.Top().curMax)
		}
		outStack.Pop()
		fmt.Fprintf(out, "%d ", maxx)

		//fmt.Println(out, inStack.items, outStack.items)
		if queue.IsEmpty() {
			break
		}

		for inStack.Size()+outStack.Size() < m {
			//fmt.Println("in")
			elem = queue.Dequeue()
			inStack.Push(elem)
		}
	}

	/*for inStack.Size() < m {
		fmt.Println("in")
		elem = queue.Dequeue()
		inStack.Push(elem)
	}

	for outStack.Size() < m {
		fmt.Println("out")
		elem = inStack.Pop().num
		outStack.Push(elem)
	}

	fmt.Println(inStack.Size(), outStack.Size())
	for inStack.Size()+outStack.Size() == m {
		fmt.Println("here")
		fmt.Println(out, inStack.items, outStack.items)
		if inStack.isEmpty() {
			maxx = outStack.Top().curMax
			outStack.Pop()
		} else if outStack.isEmpty() {
			maxx = inStack.Top().curMax
		} else {
			maxx = max(inStack.Top().curMax, outStack.Top().curMax)
			outStack.Pop()
		}
		fmt.Println(maxx)
		fmt.Println(inStack.items, outStack.items)
		if !queue.IsEmpty() {
			fmt.Println("here2")
			for inStack.Size() < m {
				fmt.Println("in")
				elem = queue.Dequeue()
				inStack.Push(elem)
			}

			for outStack.Size() < m {
				fmt.Println("out")
				elem = inStack.Pop().num
				outStack.Push(elem)
			}
		}

	}*/
}
