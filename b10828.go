package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack struct {
	items []int
}

func (s *stack) push(num int) {
	s.items = append(s.items, num)
}

func (s *stack) pop() int {
	if len(s.items) == 0 {
		return -1
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}

func (s *stack) size() int {
	return len(s.items)
}

func (s *stack) empty() int {
	if len(s.items) == 0 {
		return 1
	}
	return 0
}

func (s *stack) top() int {
	if len(s.items) == 0 {
		return -1
	}
	return s.items[len(s.items)-1]
}

func main() {
	var t int
	r := bufio.NewReader(os.Stdin)
	_, _ = fmt.Scan(&t)

	s := stack{}
	for i := 0; i < t; i++ {

		var op string
		var num int
		_, _ = fmt.Fscanln(r, &op, &num)

		switch op {
		case "push":
			s.push(num)
		case "pop":
			fmt.Println(s.pop())
		case "size":
			fmt.Println(s.size())
		case "empty":
			fmt.Println(s.empty())
		case "top":
			fmt.Println(s.top())
		}
	}
}
