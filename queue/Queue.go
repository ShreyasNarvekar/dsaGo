package queue

import "fmt"

type node struct {
	value int
	next  *node
}
type Queue struct {
	first  *node
	last   *node
	length int
}

func (Q *Queue) PrintQueue() {
	temp := Q.first
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func (Q *Queue) EnQueue(value int) {
	newNode := &node{value: value}
	if Q.length == 0 {
		Q.first = newNode
	} else {
		Q.first.next = newNode
	}
	Q.last = newNode
	Q.length++
}

func (Q *Queue) DeQueue() *node {
	if Q.length == 0 {
		return nil
	}
	temp := Q.first
	if Q.length == 1 {
		Q.first = nil
		Q.last = nil
	} else {
		Q.first = Q.first.next
		temp.next = nil
	}
	Q.length--
	return temp
}

func (Q *Queue) GetFirst() {
	fmt.Println("First element: ", Q.first.value)
}
func (Q *Queue) GetLast() {
	fmt.Println("Last element: ", Q.last.value)
}
func (Q *Queue) GetLength() {
	fmt.Println("length: ", Q.length)
}
