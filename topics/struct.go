package topics

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (user Person) printName() {
	fmt.Println(user.Name, "is ", user.Age, "years old")
}

func Structer() {
	fmt.Println("Working with Struct")

	smallPerson := Person{"John Doe", 35}
	fmt.Println(smallPerson, smallPerson.Name)

	//Working with pointers
	smallPersonCopy := &smallPerson
	fmt.Println(smallPersonCopy.Name)
	smallPersonCopy.Name = "Elizabeth Token"
	fmt.Println(smallPersonCopy.Name)

	person := Person{}
	fmt.Println(person)

	//A method can be written by using a function operating
	// a struct
	smallPerson.printName()
}
