package topics

import "fmt"

func SimpleFunction() {
	//A function that takes no argument
	fmt.Println("This is a basic function")
}

func OneArgument(amount float64) {
	fmt.Println("I was paid", amount)
}
func FunctionReturn(name string) string {
	return name
}

func SumNumbers(a int, b int) int {
	return a + b
}

func Variadic(name string) (string, string) {
	message := "I am coming"
	review := "The product is bad was his reponse says : " + name
	return message, review
}

func RestFunction(nums ...int) {
	fmt.Println("Working with a restful function")
	for _, val := range nums {
		fmt.Println(val)
	}
}
func TestFunctions() {
	defer fmt.Println("Deferred from here")
	message, _ := Variadic("Ajoke")
	fmt.Println(message)
	fmt.Println("Sum 2+ 5", SumNumbers(2, 5))
	RestFunction(2, 1, 4, 3, 6, 7)
}
