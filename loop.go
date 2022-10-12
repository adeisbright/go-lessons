package main

import "fmt"

func usingLoop(count int) {
	fmt.Println("Working with loops in Go")
	fmt.Println("---------------")
	/*
		To use the for statement,
		for initialization ; loop control , increment/decrement {
			code
		}
		Althought, the initialization and others are optional
	*/

	for i := 0; i <= count; i++ {
		fmt.Println("i is now", i)
	}
}
