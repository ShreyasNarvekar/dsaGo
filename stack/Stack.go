package stack

import (
	"errors"
	"fmt"
)

type node struct {
	Data int
	Next *node
}

type Stack struct {
	Top    *node
	height int
}

func (S *Stack) PrintStack() {
	temp := S.Top
	if S.height == 0 {
		fmt.Println("Stack is empty")
	} else {
		fmt.Println("Elements:")
		for temp != nil {
			fmt.Println(temp.Data)
			temp = temp.Next
		}
	}

}

func (S *Stack) Push(Data int) {
	newNode := &node{Data: Data}
	if S.height != 0 {
		newNode.Next = S.Top
	}
	S.Top = newNode
	S.height++
}

func (S *Stack) Pop() (int, error) {
	if S.height == 0 {
		return 0, errors.New("unable to pop. Stack is empty")
	}
	data := S.Top.Data
	temp := S.Top
	S.Top = S.Top.Next
	temp.Next = nil
	return data, nil
}

func (s *Stack) Peek() (int, error) {
	if s.height == 0 {
		return 0, errors.New("stack is empty")
	}
	return s.Top.Data, nil
}

func (S *Stack) GetHeight() {
	fmt.Println(S.height)
}

func (S *Stack) GetTop() {
	fmt.Println(S.Top.Data)
}

func (S *Stack) IsEmpty() bool {
	return S.height == 0
}
