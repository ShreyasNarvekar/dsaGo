package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastName  string
	contact   contactInfo
	//	contactInfo		we can write it like this too.
}

func main() {
	// way 1	(Preffered Way)
	// alex := person{firstname: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// way 2
	// alex := person{"Alex","Anderson"}

	// way 3 (it will put default values.	for string:""	for numeric:0	for boolean:false	 )
	// var alex person

	// By this(%+v) we can print field name and the values of field inside struct
	// fmt.Println("%+v",alex)

	// we can change/access the values by 	alex.firstname =""/fmt.pln(alex.firstName.)

	shreyas := person{
		firstname: "Shreyas",
		lastName:  "Narvekar",
		contact: contactInfo{
			email:   "abc@gmail.com",
			zipCode: 500001,
		},
	}

	shreyasPointer := &shreyas //& is passing memory address of shreyas to shreyasPointer,
	//which means shreyasPointer is a pointer pointing at shreyas of type person.

	shreyasPointer.updateName("Shree") //memoryAddress of shreyas.updateName
	shreyas.print()
}

//                    *person means accepting pointer that are pointing to person type.
func (pointerToPerson *person) updateName(newfirstname string) {

	//by putting * before pointerToPerson(having memory addr) we can Access the value in that memory addr.
	(*pointerToPerson).firstname = newfirstname
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
