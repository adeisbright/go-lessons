package main

import (
	"fmt"
	"time"
)

func timer() {
	fmt.Println("Working with Time")
	now := time.Now()
	fmt.Println(now, now.Month(), now.Year())
}
