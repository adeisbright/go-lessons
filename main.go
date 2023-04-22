package main

import (
	"fmt"
	"os"
)

type Sex struct {
	partner string
	rounds  int
}

func (sex Sex) countRounds() int {
	return sex.rounds
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Hello,World")
	//sayHello()
	// usingVariables()
	// usingSwitch(8)
	// usingLoop(10)
	// testFunctions()
	// list()
	// structer()
	// mapper()
	// timer()
	// stringer()
	//fizzBuzz(23)
	newSex := Sex{
		partner: "Blerro",
		rounds:  4,
	}

	var rounds int = newSex.countRounds()

	fmt.Println("What are we doing", rounds)
	//A Pointer is a reference to the memory address
	// to where a value is stored
	// &p : Points to the memory address of p
	// *p = value , changes p to value and dereences p

	original := "Precious Maduka"
	fake := &original
	*fake = "Girl with Big Bum Bum"
	fmt.Println(original, fake)
	//If you print a pointer , the address that the point
	// is pointed to is what will be returned

	//Working with Files

	file, err := os.ReadFile("README.md")
	checkError(err)
	fmt.Print(string(file))

	fileContent := []byte("Another content\n")
	err = os.WriteFile("file.txt", fileContent, 0666)
	checkError(err)

	/**
	os.Rename(OldName , newName) Renames a file from old to new
	os.Remove(fileName) Removes a file
	*/
}
