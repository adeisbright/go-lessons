package topics

import "fmt"

func UsingIf(num int) {
	fmt.Println("Working with if Statement")
	/*
		Using the if statement
		if [initialize the control variable] ; checking statement {
			code
		}else if condition {
			code
		}else {
			code
		}
	*/
	if a := 3; a > num {
		fmt.Println("a is greater than num")
	} else if a == num {
		fmt.Println("a is equal to num")
	} else {
		fmt.Println("a is lesser than num")
	}
}

func UsingSwitch(num int) {
	fmt.Println("Working with Switch Statement")
	/*
		Using the switch statement
		switch condition {
		case value:
			code
		case value:
			code
		default
		   code
		   breake
		}
	*/
	switch a := 3; a > num {
	case true:
		fmt.Println("a is greater than num")
	case false:
		fmt.Println("a is less than  num")
	default:
		fmt.Println("a is I cannot say")
	}
}
