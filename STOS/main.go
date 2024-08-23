package main

import (
	"fmt"
)

func main() {
	stos()
}

func stos() {
	stack := New()
	var op string
	var item int
	for {
		_, err := fmt.Scanln(&op)
		if err != nil {
			break
		}

		if op == "+" {
			fmt.Scanln(&item)
			stack.Put(item)
		} else {
			stack.Pop()
		}
	}
}

type Stack struct {
	arr []int
	ptr int
}

func New() *Stack {
	stack := new(Stack)
	stack.arr = make([]int, 10)
	stack.ptr = -1

	return stack
}

func (s *Stack) Put(item int) {
	if s.ptr+1 > 9 {
		fmt.Println(":(")
		return
	}

	s.ptr++
	s.arr[s.ptr] = item
	fmt.Println(":)")
}

func (s *Stack) Pop() {
	if s.ptr < 0 {
		fmt.Println(":(")
		return
	}

	item := s.arr[s.ptr]
	s.ptr--

	fmt.Println(item)
}
