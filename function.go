package main

import "fmt"

func simpleFunction() {
	//A function that takes no argument
	fmt.Println("This is a basic function")
}

func oneArgument(amount float64) {
	fmt.Println("I was paid", amount)
}
func functionReturn(name string) string {
	return name
}

func sumNumbers(a int, b int) int {
	return a + b
}

func variadic(name string) (string, string) {
	message := "I am coming"
	review := "The product is bad was his reponse says : " + name
	return message, review
}

func restFunction(nums ...int) {
	fmt.Println("Working with a restful function")
	for _, val := range nums {
		fmt.Println(val)
	}
}
func testFunctions() {
	defer fmt.Println("Deferred from here")
	message, _ := variadic("Ajoke")
	fmt.Println(message)
	fmt.Println("Sum 2+ 5", sumNumbers(2, 5))
	restFunction(2, 1, 4, 3, 6, 7)
}
