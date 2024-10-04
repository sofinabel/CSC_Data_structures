package main

import (
	"errors"
)

type Char struct {
	symb  rune
	index int
}

func newChar(s rune, i int) Char {
	return Char{symb: s, index: i}
}

// Stack представляет структуру стека
type Stack struct {
	items []interface{}
}

// NewStack создает новый стек
func NewStack() *Stack {
	return &Stack{
		items: []interface{}{},
	}
}

// Push добавляет элемент на вершину стека
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop удаляет и возвращает элемент с вершины стека
func (s *Stack) Pop() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, errors.New("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Peek возвращает элемент с вершины стека без его удаления
func (s *Stack) Peek() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty проверяет, пуст ли стек
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size возвращает количество элементов в стеке
func (s *Stack) Size() int {
	return len(s.items)
}

// Функция для проверки наличия символа в строке
func contains(s string, char rune) bool {
	for _, c := range s {
		if c == char {
			return true
		}
	}
	return false
}

func check(input string) int {
	stack := NewStack()
	//in := bufio.NewReader(os.Stdin)
	//out := bufio.NewWriter(os.Stdout)
	//defer out.Flush()

	//input, _ := in.ReadString('\n')
	i := 1
	var top Char
	for _, char := range input {
		if contains("({[", char) {
			stack.Push(newChar(char, i))
		} else if contains(")}]", char) {
			if stack.IsEmpty() {
				return i
			}
			topInt, _ := stack.Peek()
			top = topInt.(Char)
			if (char == ')' && top.symb == '(') || (char == ']' && top.symb == '[') || (char == '}' && top.symb == '{') {
				stack.Pop()
			} else {
				//fmt.Fprintln(out, i)
				return i
				//return
			}
		}
		i++
	}
	if stack.IsEmpty() {
		return 0
	} else {
		/*for stack.Size() > 0 {
			topInt, _ := stack.Pop()
			top = topInt.(Char)
		}*/
		topInt, _ := stack.Pop()
		top = topInt.(Char)
		return top.index
	}

	//fmt.Fprintln(out, "Success")
}
