package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stack := NewStackLink()
	stack.First = &StackLink{Item: "test"}
	stack.Size = 1

	for !stack.IsEmpty() {
		for _, v := range scanStdin() {
			tmp := string(v)
			if tmp != "-" && tmp != "\n" {
				stack.Push(tmp)
			} else if !stack.IsEmpty() {
				fmt.Println(stack.Pop())
			}
		}
		if !stack.IsEmpty() {
			fmt.Println(stack.Pop())
		}
	}
	fmt.Println("stack size left on stack is", stack.Size)
}

//StackLink 采用链表实现堆栈
type StackLink struct {
	Item        string
	First, Next *StackLink
	Size        int
}

func NewStackLink() *StackLink {
	s := &StackLink{}
	return s
}

func (s *StackLink) Push(val string) {
	oldFirst := s.First
	first := &StackLink{}
	first.Item = val
	first.Next = oldFirst
	s.First = first
	s.Size++
}

func (s *StackLink) Pop() string {
	res := s.First.Item
	s.First = s.First.Next
	s.Size--
	return res
}

//Peek 返回最近添加的元素而不弹出它
func (s *StackLink) Peek() string {
	return s.First.Item
}

func (s *StackLink) IsEmpty() bool {
	return s.First == nil
}

func scanStdin() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}
