package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	items []int
}

func newStack() *Stack {
	return &Stack{items: []int{}}
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
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

type maxStack struct {
	mainStack Stack
	maxElems  Stack
}

func newMaxStack() *maxStack {
	return &maxStack{mainStack: *newStack(), maxElems: *newStack()}
}

func (s *maxStack) Push(item int) {
	s.mainStack.Push(item)
	if s.maxElems.isEmpty() {
		s.maxElems.Push(item)
	} else {
		if s.maxElems.items[s.maxElems.Size()-1] <= item {
			s.maxElems.Push(item)
		}
	}
}

func (s *maxStack) Pop() int {
	item := s.mainStack.Pop()
	if item == s.maxElems.items[s.maxElems.Size()-1] {
		s.maxElems.Pop()
	}
	return item
}

func (s *maxStack) Output() {
	fmt.Println(s.mainStack, s.maxElems)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q, num int
	var command string
	fmt.Fscan(in, &q)
	stack := newMaxStack()
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &command)
		switch command {
		case "push":
			fmt.Fscan(in, &num)
			stack.Push(num)
		case "pop":
			stack.Pop()
		case "max":
			fmt.Fprintln(out, stack.maxElems.items[stack.maxElems.Size()-1])
		}
	}
}
