package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	text := scanStdin()
	fmt.Println(isBalance(text))
}

func isBalance(text string) bool {
	stack := NewStack()
	result := true
	for _, v := range text {
		switch v {
		case ')':
			t := stack.pop()
			if t != "(" {
				result = false
				break
			}
		case ']':
			t := stack.pop()
			if t != "[" {
				result = false
				break
			}
		case '}':
			t := stack.pop()
			if t != "{" {
				result = false
				break
			}
		case '(', '{', '[':
			stack.push(string(v))
		}
	}
	if stack.size() != 0 {
		return false
	}
	return result
}

func scanStdin() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

type Stack struct {
	a []string
}

func NewStack() *Stack {
	s := &Stack{}
	s.a = make([]string, 0)
	return s
}

func (s *Stack) isEmpty() bool {
	return len(s.a) == 0
}

func (s *Stack) size() int {
	return len(s.a)
}

func (s *Stack) push(item string) {
	s.a = append(s.a, item)
}

func (s *Stack) pop() string {
	item := s.a[len(s.a)-1]
	s.a = s.a[:len(s.a)-1]
	return item
}

//返回栈中最近添加元素而不弹出他
func (s *Stack) peek() string {
	return s.a[len(s.a)-1]
}
