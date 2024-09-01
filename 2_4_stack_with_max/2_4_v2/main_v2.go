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

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q, num int
	var command string
	fmt.Fscan(in, &q)
	stack := newStack()
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &command)
		switch command {
		case "push":
			fmt.Fscan(in, &num)
			stack.Push(num)
		case "pop":
			stack.Pop()
		case "max":
			fmt.Fprintln(out, stack.Top().curMax)
		}
	}
}
