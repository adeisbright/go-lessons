package main

import "fmt"

func list() {
	fmt.Println("Working with List")
	//An array has fixed length

	var students [3]string
	students[0] = "John Dumelo"
	students[2] = "Forcadoes"

	//You can declare and assign the values
	primes := [5]int{1, 2, 3, 5, 7}
	fmt.Println("Length of primes", len(primes))

	// A slice createa a dynamic array
	var primeSlices []int = primes[0:2]
	fmt.Println(primeSlices)

	//You can use a slice literal  to create dynamic array
	var scores = []int{10, 20, 11, 40}

	d := append(scores, 99)
	fmt.Println(d)

	m := make([]int, 0, 5) //0 is the length, 5 is the capacity
	fmt.Println(len(m), cap(m), m)
}
