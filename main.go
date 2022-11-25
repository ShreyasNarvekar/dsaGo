package main

import (
	"fmt"

	"github.com/ShreyasNarvekar/go-dsa/linkedList"
)

// "github.com/ShreyasNarvekar/go-dsa/sorting"

func main() {
	Link := new(linkedList.LinkedList) //LinkedList
	Link.Insert(10)
	Link.Insert(5)
	Link.Insert(9)
	Link.Insert(13)
	Link.Insert(22)
	Link.Insert(28)
	Link.Insert(36)
	Link.Prepend(100)
	Link.PrintList()
	fmt.Println("............................")
	fmt.Println(Link.Length())

	// var num int
	// fmt.Println("enter any number")
	// fmt.Scanln(&num)		//taking user input
	// fmt.Println("you typed:", num)

	// am := armstrong(num)
	// if am == num {
	// 	fmt.Println("is armstrong number")
	// }

	// link := new(circularll.Circularll)	//circularLL
	// link.Insert(10)
	// link.Insert(20)
	// link.PrintList()
	// link.Prepend(100)
	// link.Prepend(200)
	// link.PrintList()
	// link.RemoveFirst()
	// link.PrintList()
	// link.Removelast()
	// link.PrintList()
	// link.Reverse()
	// link.PrintList()

	// link := new(stack.Stack) //stack
	// link.Push(10)
	// link.Push(20)
	// link.Push(30)
	// link.Pop()
	// link.Push(40)
	// link.Push(50)
	// link.PrintStack()

	// array := []int{12, 6, 54, 9, 18}
	// sorting.SelectionSort(array)
	// fmt.Println(array)

	// xi := []int{1, 2, 3}
	// xii := []int{3, 4, 5}
	// copy(xi, xii)
	// fmt.Println(xi)
	// fmt.Println(len(xi))

	// fmt.Println(isValid("()"))

}

// // func armstrong
// func armstrong(num int) int {
// 	n := 0
// 	// result time
// 	temp := num
// 	for temp != 0 {
// 		temp = temp / 10
// 		n++
// 	}

// 	for temp != 0 {
// 		remainder := temp % 10
// 		result += int(math.Pow(float64(remainder), float64(n)))
// 		temp /= 10
// 	}

// 	return result

// }

// func isValid(s string) bool {
// 	isValidSym := false

// 	if len(s)%2 != 0 {
// 		return false
// 	}

// 	return isValidSym
// }
