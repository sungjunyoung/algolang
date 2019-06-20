package main

import "fmt"

type stack struct {
	node   []rune
	poperr bool
}

func (s *stack) push(n rune) {
	s.node = append(s.node, n)
}

func (s *stack) pop() {
	if len(s.node) == 0 {
		s.poperr = true
		return
	}

	s.node = s.node[0 : len(s.node)-1]
	return
}

func main() {
	var t int
	var qs string
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		s := stack{}
		succeed := false

		_, _ = fmt.Scan(&qs)
		for _, q := range qs {
			if q == '(' {
				s.push(q)
			} else if q == ')' {
				s.pop()
				if s.poperr {
					break
				}
			}

		}

		if len(s.node) == 0 && !s.poperr {
			succeed = true
		}

		if succeed {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
