package doublell

import "fmt"

type node struct {
	data int
	next *node
	prev *node
}

type DoubleLL struct {
	head   *node
	tail   *node
	length int
}

func (L *DoubleLL) Insert(data int) {
	newNode := &node{data: data}
	if L.length == 0 {
		L.head = newNode
	} else {
		L.tail.next = newNode
		newNode.prev = L.tail
	}
	L.tail = newNode
	L.length++
}

func (L *DoubleLL) Prepend(data int) {
	newNode := &node{data: data}
	if L.length == 0 {
		L.head = newNode
		L.tail = newNode
	} else {
		newNode.next = L.head
		L.head.prev = newNode
		L.head = newNode
	}
	L.length++
}

func (L *DoubleLL) PrintList() {
	temp := L.head
	for temp != nil {
		fmt.Printf("%+v ->", temp.data)
		temp = temp.next
	}
	fmt.Println()
}

func (L *DoubleLL) RemoveFirst() *node {
	if L.head == nil {
		return nil
	}
	temp := L.head
	if L.length == 1 {
		L.head = nil
		L.tail = nil
	} else {
		L.head = L.head.next
		L.head.prev = nil
		temp.next = nil
	}
	L.length--
	return temp
}

func (L *DoubleLL) Removelast() *node {
	if L.length == 0 {
		return nil
	}
	temp := L.tail
	if L.length == 1 {
		L.head = nil
		L.tail = nil
	} else {
		L.tail = L.tail.prev
		L.tail.next = nil
		temp.prev = nil
	}
	L.length--
	return temp
}

func (L *DoubleLL) GetNode(index int) *node {
	if index < 0 || index >= L.length {
		return nil
	}
	temp := L.head
	if index < L.length/2 {
		for i := 0; i < index; i++ {
			temp = temp.next
		}
	} else {
		temp = L.tail
		for i := L.length - 1; i > index; i-- {
			temp = temp.prev
		}
	}
	return temp
}

func (L *DoubleLL) SetNode(index, value int) bool {
	temp := L.GetNode(index)
	if temp != nil {
		temp.data = value
		return true
	}
	return false
}

func (L *DoubleLL) InsertAt(index, value int) bool {
	if index < 0 || index > L.length {
		return false
	}
	if index == 0 {
		L.Prepend(value)
		return true
	}
	if index == L.length {
		L.Insert(value)
		return true
	}
	newNode := &node{data: value}
	temp := L.GetNode(index)
	newNode.prev = temp.prev
	newNode.next = temp
	temp.prev.next = newNode
	temp.prev = newNode
	L.length++
	return true
}

func (L *DoubleLL) RemoveAt(index int) *node {
	if index < 0 || index >= L.length {
		return nil
	}
	if index == 0 {
		return L.RemoveFirst()
	}
	if index == L.length-1 {
		L.Removelast()
	}
	temp := L.GetNode(index)
	temp.prev.next = temp.next
	temp.next.prev = temp.prev
	temp.next = nil
	temp.prev = nil
	L.length--
	return temp
}





























