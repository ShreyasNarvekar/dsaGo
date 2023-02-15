package main

import "fmt"

// interface
type bot interface {
	getGreeting() string //any type having this mwthod returning string are now also bot type.
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

//if we are not using variable of reciever type than we can delete it
//from..func(eb englishBot)getGreeting() we can do..
// 		func(englishBot)getGreeting()
func (englishBot) getGreeting() string {
	return "Hi There"
}
func (spanishBot) getGreeting() string {
	return "Hola!"
}

//because english/spanishbot both are now bot type,they can now use func accepting bot arguments.
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
