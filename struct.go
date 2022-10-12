package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (user Person) printName() {
	fmt.Println(user.name, "is ", user.age, "years old")
}

func structer() {
	fmt.Println("Working with Struct")

	smallPerson := Person{"John Doe", 35}
	fmt.Println(smallPerson, smallPerson.name)

	//Working with pointers
	smallPersonCopy := &smallPerson
	fmt.Println(smallPersonCopy.name)
	smallPersonCopy.name = "Elizabeth Token"
	fmt.Println(smallPersonCopy.name)

	person := Person{}
	fmt.Println(person)

	//A method can be written by using a function operating
	// a struct
	smallPerson.printName()
}
