package main

import "fmt"

func usingVariables() {
	//A String
	var name string = "Adeleke Bright"
	fmt.Println(name)
	//integer
	var age int = 18
	//floating point
	var laptopPrice float64 = 35.29
	//boolean
	var isOld bool = false

	fmt.Println("The age is", age)
	fmt.Println(laptopPrice, isOld)

	//This cannot be done outside a function
	// you cannot also use it for a variable that has
	// been declared before
	language := "Go Programming Language"
	fmt.Println("I am building with", language)

	var (
		country       = "Nigeria"
		independence  = 1960
		usesDemocracy = true
	)
	fmt.Println(country, "has democracy?", usesDemocracy, independence)
}
