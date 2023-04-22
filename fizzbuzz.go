package main

import "fmt"

func fizzBuzz(number int) {
	fmt.Println("Working with Fizzbuzz")
	for i := 0; i <= number; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(number)
		}
	}
}
